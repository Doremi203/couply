package user_service

import (
	"context"
	"github.com/Doremi203/Couply/backend/internal/dto"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateUserV1(ctx context.Context, in *desc.UpdateUserV1Request) (*desc.UpdateUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.UpdateUser(ctx, dto.PBToUpdateUserRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.UpdateUserResponseToPB(response), nil
}
