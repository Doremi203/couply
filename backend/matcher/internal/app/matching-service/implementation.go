package matching_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type matchingServiceUseCase interface {
	LikeUser(ctx context.Context, in *dto.LikeUserV1Request) (*dto.LikeUserV1Response, error)
	DislikeUser(ctx context.Context, in *dto.DislikeUserV1Request) (*dto.DislikeUserV1Response, error)
	DeleteMatch(ctx context.Context, in *dto.DeleteMatchV1Request) (*dto.DeleteMatchV1Response, error)
	FetchMatchesUserIDs(ctx context.Context, in *dto.FetchMatchesUserIDsV1Request) (*dto.FetchMatchesUserIDsV1Response, error)
	FetchOutgoingLikes(ctx context.Context, in *dto.FetchOutgoingLikesV1Request) (*dto.FetchOutgoingLikesV1Response, error)
	FetchIncomingLikes(ctx context.Context, in *dto.FetchIncomingLikesV1Request) (*dto.FetchIncomingLikesV1Response, error)
}

type Implementation struct {
	desc.UnimplementedMatchingServiceServer
	usecase matchingServiceUseCase
}

func NewImplementation(
	usecase matchingServiceUseCase,
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
	return desc.RegisterMatchingServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (i *Implementation) RegisterToServer(gRPC *grpc.Server) {
	desc.RegisterMatchingServiceServer(gRPC, i)
}
