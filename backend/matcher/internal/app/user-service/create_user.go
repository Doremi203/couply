package user_service

import (
	"context"
	"github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateUserV1(ctx context.Context, in *desc.CreateUserV1Request) (*desc.CreateUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.CreateUser(ctx, user_service.PBToCreateUserRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user_service.CreateUserResponseToPB(response), nil
}
