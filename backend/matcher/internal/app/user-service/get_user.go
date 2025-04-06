package user_service

import (
	"context"
	user_service "github.com/Doremi203/Couply/backend/internal/dto/user-service"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUserV1(ctx context.Context, in *desc.GetUserV1Request) (*desc.GetUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetUser(ctx, user_service.PBToGetUserRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user_service.GetUserResponseToPB(response), nil
}
