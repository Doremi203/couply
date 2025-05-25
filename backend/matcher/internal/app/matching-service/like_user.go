package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/matching-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) LikeUserV1(ctx context.Context, in *desc.LikeUserV1Request) (*desc.LikeUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToLikeUserRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.LikeUser(ctx, req)
	switch {
	case errors.Is(err, matching.ErrLikeNotFound):
		return nil, status.Error(codes.NotFound, matching.ErrLikeNotFound.Error())
	case errors.Is(err, user.ErrUserDoesntExist):
		return nil, status.Error(codes.FailedPrecondition, user.ErrUserDoesntExist.Error())
	case errors.Is(err, matching.ErrMatchAlreadyExists):
		return nil, status.Error(codes.FailedPrecondition, matching.ErrMatchAlreadyExists.Error())
	case err != nil:
		return nil, err
	}

	return dto.LikeUserResponseToPB(response), nil
}
