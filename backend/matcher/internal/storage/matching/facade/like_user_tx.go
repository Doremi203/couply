package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/matching/postgres"
)

func (f *StorageFacadeMatching) LikeUserTx(ctx context.Context, like *matching.Like) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		gottenLike, err := f.storage.GetLike(ctxTx, postgres.GetLikeOptions{
			SenderID:   like.SenderID,
			ReceiverID: like.ReceiverID,
			IsWaiting:  true,
		})
		if err != nil {
			if !errors.Is(err, matching.ErrLikeNotFound) {
				return errors.Wrap(err, "storage.GetLike")
			}
		}

		if gottenLike != nil {
			return matching.ErrWaitingLikeAlreadyExists
		}

		match, err := f.storage.GetMatch(ctx, postgres.GetMatchOptions{FirstUserID: like.SenderID, SecondUserID: like.ReceiverID})
		if err != nil {
			if !errors.Is(err, matching.ErrMatchNotFound) {
				return errors.Wrap(err, "storage.GetMatch")
			}
		}

		if match != nil {
			return matching.ErrMatchExistsBetweenTheseUsers
		}

		if err = f.storage.CreateLike(ctxTx, like); err != nil {
			return errors.Wrap(err, "storage.CreateLike")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
