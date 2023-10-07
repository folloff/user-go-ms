package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"auth-go/internal/data"
	"auth-go/pkg/logger"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8080"

var reconnectionCounts int64

type Config struct {
	DB     *sqlx.DB
	Models data.Models
}

func main() {
	startMessage := fmt.Sprintf("Starting API server on port %s", webPort)
	logger.Info(startMessage)

	// TODO: connect to DB
	conn := connectToDB()
	if conn == nil {
		logger.Error("Cannot connect to DB")
	}

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Check connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sqlx.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/folloff-auth?sslmode=disable"

	for {
		connection, err := openDB(dsn)

		if err != nil {
			logger.Error("Cannot connect to DB: " + err.Error())
			reconnectionCounts++
		} else {
			logger.Info("Connected to DB")
			return connection
		}

		if reconnectionCounts > 10 {
			logger.Error("Cannot connect to DB after 10 ettempts. Last error : " + err.Error())
			return nil
		}
		logger.Info("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
