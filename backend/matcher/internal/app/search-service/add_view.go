package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"

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
	switch {
	case errors.Is(err, user.ErrUserDoesntExist):
		return nil, status.Error(codes.FailedPrecondition, user.ErrUserDoesntExist.Error())
	case err != nil:
		return nil, err
	}

	return dto.AddViewResponseToPB(response), nil
}
