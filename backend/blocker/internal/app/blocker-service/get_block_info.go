package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"

	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetBlockInfoV1(ctx context.Context, in *desc.GetBlockInfoV1Request) (*desc.GetBlockInfoV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.GetBlockInfo(ctx, dto.PBToGetBlockInfoRequest(in))
	switch {
	case errors.Is(err, blocker.ErrUserBlockNotFound):
		return nil, status.Error(codes.NotFound, blocker.ErrUserBlockNotFound.Error())
	case errors.Is(err, blocker.ErrUserBlockReasonsNotFound):
		return nil, status.Error(codes.NotFound, blocker.ErrUserBlockReasonsNotFound.Error())
	case err != nil:
		return nil, err
	}

	return dto.GetBlockInfoResponseToPB(response), nil
}
