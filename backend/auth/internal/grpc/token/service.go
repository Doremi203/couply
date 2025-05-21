package token

import (
	"context"

	tokenPb "github.com/Doremi203/couply/backend/auth/gen/api/token"
	tokenUC "github.com/Doremi203/couply/backend/auth/internal/usecase/token"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewGRPCService(
	logger log.Logger,
	useCase tokenUC.UseCase,
) *grpcService {
	return &grpcService{
		useCase: useCase,
		logger:  logger,
	}
}

type grpcService struct {
	useCase tokenUC.UseCase

	logger log.Logger
	tokenPb.UnimplementedTokenProviderServer
}

func (s *grpcService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return tokenPb.RegisterTokenProviderHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *grpcService) RegisterToServer(gRPC *grpc.Server) {
	tokenPb.RegisterTokenProviderServer(gRPC, s)
}
