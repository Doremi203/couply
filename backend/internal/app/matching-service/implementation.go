package matching_service

import (
	"context"
	dto "github.com/Doremi203/Couply/backend/internal/dto/matching-service"
	desc "github.com/Doremi203/Couply/backend/pkg/matching-service/v1"
)

type matchingServiceUseCase interface {
	CreateMatch(ctx context.Context, in *dto.CreateMatchV1Request) (*dto.CreateMatchV1Response, error)
	UpdateMatch(ctx context.Context, in *dto.UpdateMatchV1Request) (*dto.UpdateMatchV1Response, error)
	FetchMatches(ctx context.Context, in *dto.FetchMatchesV1Request) (*dto.FetchMatchesV1Response, error)
	FetchIncomingMatches(ctx context.Context, in *dto.FetchIncomingMatchesV1Request) (*dto.FetchIncomingMatchesV1Response, error)
	FetchOutgoingMatches(ctx context.Context, in *dto.FetchOutgoingMatchesV1Request) (*dto.FetchOutgoingMatchesV1Response, error)
}

type Implementation struct {
	desc.UnimplementedMatchingServiceServer
	usecase matchingServiceUseCase
}

func NewImplementation(usecase matchingServiceUseCase) *Implementation {
	return &Implementation{usecase: usecase}
}
