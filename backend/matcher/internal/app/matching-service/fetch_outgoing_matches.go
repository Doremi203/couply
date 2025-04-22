package matching_service

import (
	"context"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) FetchOutgoingMatchesV1(ctx context.Context, in *desc.FetchOutgoingMatchesV1Request) (*desc.FetchOutgoingMatchesV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.FetchOutgoingMatches(ctx, dto.PBToFetchOutgoingMatchesRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FetchOutgoingMatchesResponseToPB(response), nil
}
