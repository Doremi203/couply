package blocker_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type blockerServiceUseCase interface {
	ReportUser(ctx context.Context, in *dto.ReportUserV1Request) (*dto.ReportUserV1Response, error)
	GetBlockInfo(ctx context.Context, in *dto.GetBlockInfoV1Request) (*dto.GetBlockInfoV1Response, error)
}

type Implementation struct {
	desc.UnimplementedBlockerServiceServer
	usecase blockerServiceUseCase
}

func NewImplementation(
	usecase blockerServiceUseCase,
) *Implementation {
	return &Implementation{
		usecase: usecase,
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
