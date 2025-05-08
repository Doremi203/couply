package matching_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) FetchIncomingLikesV1(ctx context.Context, in *desc.FetchIncomingLikesV1Request) (*desc.FetchIncomingLikesV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := dto.PBToFetchIncomingLikesRequest(in)

	response, err := i.usecase.FetchIncomingLikes(ctx, req)
	if err != nil {
		i.logger.Error(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return dto.FetchIncomingLikesResponseToPB(response), nil
}
