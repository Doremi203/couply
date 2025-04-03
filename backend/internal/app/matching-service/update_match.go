package matching_service

import (
	"context"
	dto "github.com/Doremi203/Couply/backend/internal/dto/matching-service"
	desc "github.com/Doremi203/Couply/backend/pkg/matching-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateMatchV1(ctx context.Context, in *desc.UpdateMatchV1Request) (*desc.UpdateMatchV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.UpdateMatch(ctx, dto.PBToUpdateMatchRequest(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.UpdateMatchResponseToPB(response), nil
}
