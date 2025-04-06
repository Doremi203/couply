package user_service

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/dto/user-service"

	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteUserV1(ctx context.Context, in *desc.DeleteUserV1Request) (*desc.DeleteUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.DeleteUser(ctx, user_service.PBToDeleteUserRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user_service.DeleteUserResponseToPB(response), nil
}
