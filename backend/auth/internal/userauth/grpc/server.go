package grpc

import (
	"auth/proto/userauth"
	"context"
	"google.golang.org/grpc"
)

type server struct {
	userauth.UnimplementedUserAuthServer
}

func RegisterServer(gRPC *grpc.Server) {
	userauth.RegisterUserAuthServer(gRPC, &server{})
}

func (s *server) Register(
	ctx context.Context,
	req *userauth.RegisterRequest,
) (*userauth.RegisterResponse, error) {
	return nil, nil
}
