package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteUserV1(ctx context.Context, in *desc.DeleteUserV1Request) (*desc.DeleteUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.DeleteUser(ctx, dto.PBToDeleteUserRequest(in))
	switch {
	case errors.Is(err, user.ErrUserNotFound):
		return nil, status.Error(codes.NotFound, user.ErrUserNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.DeleteUserResponseToPB(response), nil
}
