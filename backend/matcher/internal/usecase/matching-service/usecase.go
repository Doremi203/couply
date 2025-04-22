package matching_service

import (
	"context"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

type matchingStorageFacade interface {
	CreateMatchTx(ctx context.Context, match *matching.Match) (*matching.Match, error)
	UpdateMatchTx(ctx context.Context, match *matching.Match) (*matching.Match, error)
	FetchMatchesTx(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error)
	FetchIncomingMatchesTx(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error)
	FetchOutgoingMatchesTx(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error)
}

type UseCase struct {
	matchingStorageFacade matchingStorageFacade
}

func NewUseCase(matchingStorageFacade matchingStorageFacade) *UseCase {
	return &UseCase{matchingStorageFacade: matchingStorageFacade}
}
