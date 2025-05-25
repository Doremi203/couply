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

func (i *Implementation) DislikeUserV1(ctx context.Context, in *desc.DislikeUserV1Request) (*desc.DislikeUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToDislikeUserRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.DislikeUser(ctx, req)
	switch {
	case errors.Is(err, matching.ErrLikeNotFound):
		return nil, status.Error(codes.NotFound, matching.ErrLikeNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.DislikeUserResponseToPB(response), nil
}
