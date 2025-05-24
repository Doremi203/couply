package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

func (f *StorageFacadeMatching) GetLikeTx(ctx context.Context, senderID, receiverID uuid.UUID) (*matching.Like, error) {
	var like *matching.Like
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		like, err = f.storage.GetLike(ctx, senderID, receiverID)
		if err != nil {
			return err
		}

		return err
	})

	return like, err
}
