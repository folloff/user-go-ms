package cmd

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"

	"github.com/brianvoe/gofakeit"

	desc "github.com/folloff/auth-go-ms/pkg/auth_v1"
)

const grpcPort = 50055

type server struct {
	desc.UnimplementedUserV1Server
}

// Get ...
func (s *server) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	log.Printf("Note id: %d", req.GetPublicId())

	return &desc.GetUserResponse{
		Data: &desc.UserPublicData{
			PublicId:  req.GetPublicId(),
			Name:      gofakeit.BeerName(),
			Email:     gofakeit.Email(),
			Role:      0,
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {

}
