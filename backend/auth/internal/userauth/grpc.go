package userauth

import (
	"auth/proto/userauth"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func NewGRPCService() *gRPCService {
	return &gRPCService{}
}

type gRPCService struct {
	userauth.UnimplementedUserAuthServer
}

func (s *gRPCService) Register(gRPC *grpc.Server) {
	userauth.RegisterUserAuthServer(gRPC, s)
}

func (s *gRPCService) RegisterUser(
	ctx context.Context,
	req *userauth.RegisterRequest,
) (*userauth.RegisterResponse, error) {
	fmt.Println("gRPC RegisterUser")
	return nil, nil
}
