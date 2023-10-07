package cmd

import (
	desc "github.com/folloff/auth-go-ms/pkg/auth_v1"
)

const grpcPort = 50055

type server struct {
	desc.UnimplementedAuthV1Server
}

func main() {

}
