package facade

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/matching/postgres"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

func (f *StorageFacadeMatching) GetLikeTx(ctx context.Context, senderID, receiverID uuid.UUID) (*matching.Like, error) {
	like, err := f.storage.GetLike(ctx, postgres.GetLikeOptions{
		SenderID:   senderID,
		ReceiverID: receiverID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "storage.GetLike")
	}

	return like, nil
}
