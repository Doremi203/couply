package search_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) AddViewV1(ctx context.Context, in *desc.AddViewV1Request) (*desc.AddViewV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToAddViewRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.AddView(ctx, req)
	if err != nil {
		i.logger.Error(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return dto.AddViewResponseToPB(response), nil
}
