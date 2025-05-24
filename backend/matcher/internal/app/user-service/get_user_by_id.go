package user_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/user-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUserByIDV1(ctx context.Context, in *desc.GetUserByIDV1Request) (*desc.GetUserByIDV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToGetUserByIDRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetUserByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return dto.GetUserByIDResponseToPB(response), nil
}
