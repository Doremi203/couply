package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/log"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type searchServiceUseCase interface {
	CreateFilter(ctx context.Context, in *dto.CreateFilterV1Request) (*dto.CreateFilterV1Response, error)
	UpdateFilter(ctx context.Context, in *dto.UpdateFilterV1Request) (*dto.UpdateFilterV1Response, error)
	GetFilter(ctx context.Context, in *dto.GetFilterV1Request) (*dto.GetFilterV1Response, error)
	SearchUsers(ctx context.Context, in *dto.SearchUsersV1Request) (*dto.SearchUsersV1Response, error)
	AddView(ctx context.Context, in *dto.AddViewV1Request) (*dto.AddViewV1Response, error)
}

type Implementation struct {
	desc.UnimplementedSearchServiceServer
	usecase           searchServiceUseCase
	photoURLGenerator user.PhotoURLGenerator
	logger            log.Logger
}

func NewImplementation(
	logger log.Logger,
	usecase searchServiceUseCase,
	photoURLGenerator user.PhotoURLGenerator,
) *Implementation {
	return &Implementation{
		logger:            logger,
		usecase:           usecase,
		photoURLGenerator: photoURLGenerator,
	}
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
