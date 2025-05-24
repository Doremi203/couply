package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

type matchingStorageFacade interface {
	LikeUserTx(ctx context.Context, like *matching.Like) error
	GetLikeTx(ctx context.Context, senderID, receiverID uuid.UUID) (*matching.Like, error)
	UpdateLikeTx(ctx context.Context, like *matching.Like) error
	DeleteMatchTx(ctx context.Context, userID, targetUserID uuid.UUID) error
	FetchMatchesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Match, error)
	FetchOutgoingLikesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Like, error)
	FetchIncomingLikesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Like, error)
	HandleMutualLikeTx(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*matching.Match, error)
}

type UseCase struct {
	matchingStorageFacade matchingStorageFacade
}

func NewUseCase(matchingStorageFacade matchingStorageFacade) *UseCase {
	return &UseCase{matchingStorageFacade: matchingStorageFacade}
}
