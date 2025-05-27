//go:generate mockgen -source=usecase.go -destination=../../mocks/usecase/matching/facade_mock.go -typed

package matching_service

import (
	"context"

	"github.com/Doremi203/couply/backend/common/libs/sqs"
	"github.com/Doremi203/couply/backend/matcher/gen/api/messages"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

type matchingStorageFacade interface {
	matchingStorageSetterFacade
	matchingStorageGetterFacade
}

type matchingStorageSetterFacade interface {
	LikeUserTx(ctx context.Context, like *matching.Like) error
	UpdateLikeTx(ctx context.Context, like *matching.Like) error
	DeleteMatchTx(ctx context.Context, userID, targetUserID uuid.UUID) error
	HandleMutualLikeTx(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*matching.Match, error)
}

type matchingStorageGetterFacade interface {
	GetLikeTx(ctx context.Context, senderID, receiverID uuid.UUID) (*matching.Like, error)
	GetWaitingLikeTx(ctx context.Context, senderID, receiverID uuid.UUID) (*matching.Like, error)
	FetchMatchesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Match, error)
	FetchOutgoingLikesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Like, error)
	FetchIncomingLikesTx(ctx context.Context, userID uuid.UUID, limit, offset uint64) ([]*matching.Like, error)
}

type UseCase struct {
	matchingStorageFacade matchingStorageFacade
	sqsClient             sqs.ClientWriter[*messages.MatcherEvent]
}

func NewUseCase(
	matchingStorageFacade matchingStorageFacade,
	client sqs.ClientWriter[*messages.MatcherEvent],
) *UseCase {
	return &UseCase{
		matchingStorageFacade: matchingStorageFacade,
		sqsClient:             client,
	}
}
