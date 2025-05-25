package matching_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) FetchMatchesUserIDsV1(ctx context.Context, in *desc.FetchMatchesUserIDsV1Request) (*desc.FetchMatchesUserIDsV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.FetchMatchesUserIDs(ctx, dto.PBToFetchMatchesUserIDsRequest(in))
	if err != nil {
		return nil, err
	}

	return dto.FetchMatchesUserIDsResponseToPB(response), nil
}
