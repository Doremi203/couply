package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteMatchV1(ctx context.Context, in *desc.DeleteMatchV1Request) (*desc.DeleteMatchV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToDeleteMatchRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.DeleteMatch(ctx, req)
	switch {
	case errors.Is(err, matching.ErrMatchNotFound):
		return nil, status.Error(codes.NotFound, matching.ErrMatchNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.DeleteMatchResponseToPB(response), nil
}
