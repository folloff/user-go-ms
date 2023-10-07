package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "timestamp"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.StacktraceKey = ""
	config.EncoderConfig = encodeConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	// log, err = zap.NewProduction(zap.AddCallerSkip(1)) << It was default

	if err != nil {
		panic(err)
	}
}

type Interface interface {
	Info(message string, args ...interface{})
	Debug(message string, args ...interface{})
//	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
//	Fatal(message interface{}, args ...interface{})
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
