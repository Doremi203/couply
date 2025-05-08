package user_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	user_service "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateUserV1(ctx context.Context, in *desc.UpdateUserV1Request) (*desc.UpdateUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.UpdateUser(ctx, user_service.PBToUpdateUserRequest(in))
	if err != nil {
		i.logger.Error(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return user_service.UpdateUserResponseToPB(response), nil
}
