package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	dto "github.com/Doremi203/couply/backend/matcher/internal/dto/search-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetFilterV1(ctx context.Context, in *desc.GetFilterV1Request) (*desc.GetFilterV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetFilter(ctx, dto.PBToGetFilterRequest(in))
	switch {
	case errors.Is(err, search.ErrFilterNotFound):
		return nil, status.Error(codes.NotFound, search.ErrFilterNotFound.Error())
	case errors.Is(err, interest.ErrInterestsNotFound):
		return nil, status.Error(codes.NotFound, interest.ErrInterestsNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.GetFilterResponseToPB(response), nil
}
