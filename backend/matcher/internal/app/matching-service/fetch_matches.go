package matching_service

import (
	"context"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) FetchMatchesV1(ctx context.Context, in *desc.FetchMatchesV1Request) (*desc.FetchMatchesV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.FetchMatches(ctx, dto.PBToFetchMatchesRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FetchMatchesResponseToPB(response), nil
}
