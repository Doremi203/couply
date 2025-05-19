package blocker_service

import (
	"context"

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
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.GetBlockInfoResponseToPB(response), nil
}
