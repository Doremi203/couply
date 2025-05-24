package facade

import (
	"context"

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
			return err
		}

		newLike := matching.NewLike(userID, targetUserID, message, matching.StatusAccepted)
		err = f.storage.AddLike(ctx, newLike)
		if err != nil {
			return err
		}

		newMatch = matching.NewMatch(userID, targetUserID)
		err = f.storage.AddMatch(ctx, newMatch)
		if err != nil {
			return err
		}

		return err
	})

	return newMatch, err
}
