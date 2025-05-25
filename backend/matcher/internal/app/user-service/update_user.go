package user_service

import (
	"context"

	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateUserV1(ctx context.Context, in *desc.UpdateUserV1Request) (*desc.UpdateUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.UpdateUser(ctx, dto.PBToUpdateUserRequest(in))
	if err != nil {
		return nil, err
	}

	return dto.UpdateUserResponseToPB(response), nil
}
