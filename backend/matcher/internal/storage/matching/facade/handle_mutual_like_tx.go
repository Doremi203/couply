package facade

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/google/uuid"
)

func (f *StorageFacadeMatching) HandleMutualLikeTx(ctx context.Context, userID, targetUserID uuid.UUID, message string) (*matching.Match, error) {
	var newMatch *matching.Match
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		updatedLike := matching.NewLike(targetUserID, userID, message, matching.StatusAccepted)
		err = f.storage.UpdateLike(ctx, updatedLike)
		if err != nil {
			return errors.Wrap(err, "storage.UpdateLike")
		}

		newLike := matching.NewLike(userID, targetUserID, message, matching.StatusAccepted)
		err = f.storage.CreateLike(ctx, newLike)
		if err != nil {
			return errors.Wrap(err, "storage.CreateLike")
		}

		newMatch = matching.NewMatch(userID, targetUserID)
		err = f.storage.CreateMatch(ctx, newMatch)
		if err != nil {
			return errors.Wrap(err, "storage.CreateMatch")
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return newMatch, nil
}
