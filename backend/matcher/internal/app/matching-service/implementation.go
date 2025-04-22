package matching_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type matchingServiceUseCase interface {
	CreateMatch(ctx context.Context, in *dto.CreateMatchV1Request) (*dto.CreateMatchV1Response, error)
	UpdateMatch(ctx context.Context, in *dto.UpdateMatchV1Request) (*dto.UpdateMatchV1Response, error)
	FetchMatches(ctx context.Context, in *dto.FetchMatchesV1Request) (*dto.FetchMatchesV1Response, error)
	FetchIncomingMatches(ctx context.Context, in *dto.FetchIncomingMatchesV1Request) (*dto.FetchIncomingMatchesV1Response, error)
	FetchOutgoingMatches(ctx context.Context, in *dto.FetchOutgoingMatchesV1Request) (*dto.FetchOutgoingMatchesV1Response, error)
}

type Implementation struct {
	desc.UnimplementedMatchingServiceServer
	usecase matchingServiceUseCase
}

func NewImplementation(usecase matchingServiceUseCase) *Implementation {
	return &Implementation{usecase: usecase}
}

func (i *Implementation) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return desc.RegisterMatchingServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (i *Implementation) RegisterToServer(gRPC *grpc.Server) {
	desc.RegisterMatchingServiceServer(gRPC, i)
}
