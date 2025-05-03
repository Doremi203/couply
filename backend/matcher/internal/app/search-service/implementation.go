package search_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type searchServiceUseCase interface {
	CreateFilter(ctx context.Context, in *dto.CreateFilterV1Request) (*dto.CreateFilterV1Response, error)
	UpdateFilter(ctx context.Context, in *dto.UpdateFilterV1Request) (*dto.UpdateFilterV1Response, error)
	GetFilter(ctx context.Context, in *dto.GetFilterV1Request) (*dto.GetFilterV1Response, error)
}

type Implementation struct {
	desc.UnimplementedSearchServiceServer
	usecase searchServiceUseCase
}

func NewImplementation(usecase searchServiceUseCase) *Implementation {
	return &Implementation{usecase: usecase}
}

func (i *Implementation) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return desc.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (i *Implementation) RegisterToServer(gRPC *grpc.Server) {
	desc.RegisterSearchServiceServer(gRPC, i)
}
