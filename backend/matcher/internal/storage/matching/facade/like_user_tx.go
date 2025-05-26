package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/matching/postgres"
)

func (f *StorageFacadeMatching) LikeUserTx(ctx context.Context, like *matching.Like) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if _, err := f.storage.GetLike(ctx, postgres.GetLikeOptions{
			SenderID:   like.SenderID,
			ReceiverID: like.ReceiverID,
			IsWaiting:  true,
		}); err != nil {
			return errors.Wrap(matching.ErrWaitingLikeAlreadyExists, "storage.GetLike")
		}

		if err := f.storage.CreateLike(ctx, like); err != nil {
			return errors.Wrap(err, "storage.CreateLike")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
