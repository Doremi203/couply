package search_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeSearch) SearchUsersTx(ctx context.Context, userID uuid.UUID, page, limit uint64) ([]*user.User, map[uuid.UUID]float64, error) {
	var (
		users     []*user.User
		distances map[uuid.UUID]float64
		err       error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		fil, err := f.searchStorage.GetFilter(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: get user failed: %w", err)
		}

		filI, err := f.searchStorage.GetFilterInterests(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: get interests failed: %w", err)
		}

		curUser, err := f.userStorage.GetUser(ctx, userID)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: get user failed: %w", err)
		}

		users, distances, err = f.searchStorage.SearchUsers(ctx, fil, filI, curUser.Latitude, curUser.Longitude, page, limit)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: search failed: %w", err)
		}

		for _, u := range users {
			userInterest, err := f.userStorage.GetInterests(ctx, u.ID)
			if err != nil {
				return fmt.Errorf("SearchUsersTx: get interests failed: %w", err)
			}

			u.Interest = userInterest
		}

		for _, u := range users {
			userPhotos, err := f.userStorage.GetPhotos(ctx, u.ID)
			if err != nil {
				return fmt.Errorf("SearchUsersTx: get photos failed: %w", err)
			}

			u.Photos = userPhotos
		}

		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("SearchUsersTx: user transaction failed: %w", err)
	}
	return users, distances, nil
}
