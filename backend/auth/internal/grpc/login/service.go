package login

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/gen/api/login"
	loginUC "github.com/Doremi203/couply/backend/auth/internal/usecase/login"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewGRPCService(
	loginUseCase loginUC.UseCase,
	logger log.Logger,
) *grpcService {
	return &grpcService{
		loginUseCase: loginUseCase,
		logger:       logger,
	}
}

type grpcService struct {
	loginUseCase loginUC.UseCase

	logger log.Logger
	login.UnimplementedLoginServer
}

func (s *grpcService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return login.RegisterLoginHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *grpcService) RegisterToServer(gRPC *grpc.Server) {
	login.RegisterLoginServer(gRPC, s)
}
