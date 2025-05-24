package user_service

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

func (f *StorageFacadeUser) GetUsersTx(ctx context.Context, userIDs []uuid.UUID) ([]*user.User, error) {
	var users []*user.User

	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		var err error

		users, err = f.storage.GetUsers(ctxTx, userIDs)
		if err != nil {
			return fmt.Errorf("GetUsersTx: get users failed: %w", err)
		}

		photosMap, err := f.storage.GetPhotosForUsers(ctxTx, userIDs)
		if err != nil {
			return fmt.Errorf("GetUsersTx: get photos failed: %w", err)
		}

		interestsMap, err := f.storage.GetInterestsForUsers(ctxTx, userIDs)
		if err != nil {
			return fmt.Errorf("GetUsersTx: get interests failed: %w", err)
		}

		for _, u := range users {
			if photos, exists := photosMap[u.ID]; exists {
				u.Photos = photos
			}

			if interest, exists := interestsMap[u.ID]; exists {
				u.Interest = interest
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("GetUsersTx: users transaction failed: %w", err)
	}
	return users, nil
}
