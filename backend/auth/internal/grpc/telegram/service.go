package telegram

import (
	"context"

	telegramPb "github.com/Doremi203/couply/backend/auth/gen/api/telegram"
	telegramUC "github.com/Doremi203/couply/backend/auth/internal/usecase/telegram"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewGRPCService(
	logger log.Logger,
	useCase telegramUC.UseCase,
) *grpcService {
	return &grpcService{
		useCase: useCase,
		logger:  logger,
	}
}

type grpcService struct {
	useCase telegramUC.UseCase

	logger log.Logger
	telegramPb.UnimplementedTelegramDataServer
}

func (s *grpcService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return telegramPb.RegisterTelegramDataHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *grpcService) RegisterToServer(gRPC *grpc.Server) {
	telegramPb.RegisterTelegramDataServer(gRPC, s)
}
