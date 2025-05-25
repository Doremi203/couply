package user_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateUserV1(ctx context.Context, in *desc.CreateUserV1Request) (*desc.CreateUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := dto.PBToCreateUserRequest(in)

	response, err := i.usecase.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return dto.CreateUserResponseToPB(response), nil
}
