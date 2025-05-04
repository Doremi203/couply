package search_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SearchUsersV1(ctx context.Context, in *desc.SearchUsersV1Request) (*desc.SearchUsersV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.SearchUsers(ctx, dto.PBToSearchUsersRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.SearchUsersResponseToPB(response), nil
}
