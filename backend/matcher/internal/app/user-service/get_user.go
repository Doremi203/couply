package user_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUserV1(ctx context.Context, in *desc.GetUserV1Request) (*desc.GetUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetUser(ctx, dto.PBToGetUserRequest(in))
	if err != nil {
		return nil, err
	}

	return dto.GetUserResponseToPB(response), nil
}
