package user_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUsersV1(ctx context.Context, in *desc.GetUsersV1Request) (*desc.GetUsersV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dtoReq, err := dto.PBToGetUsersRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetUsers(ctx, dtoReq)
	if err != nil {
		return nil, err
	}

	return dto.GetUsersResponseToPB(response), nil
}
