package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/log"
	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type blockerServiceUseCase interface {
	ReportUser(ctx context.Context, request *dto.ReportUserRequest) (*dto.ReportUserResponse, error)
}

type Implementation struct {
	desc.UnimplementedBlockerServiceServer
	usecase blockerServiceUseCase
	logger  log.Logger
}

func NewImplementation(
	usecase blockerServiceUseCase,
	logger log.Logger,
) *Implementation {
	return &Implementation{
		usecase: usecase,
		logger:  logger,
	}
}

func (i *Implementation) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return desc.RegisterBlockerServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (i *Implementation) RegisterToServer(gRPC *grpc.Server) {
	desc.RegisterBlockerServiceServer(gRPC, i)
}
