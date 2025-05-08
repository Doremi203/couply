package search_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeSearch) SearchUsersTx(ctx context.Context, userID uuid.UUID, page, limit uint64) ([]*user.User, error) {
	var (
		fil   *search.Filter
		filI  *interest.Interest
		users []*user.User
		err   error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		fil, err = f.searchStorage.GetFilter(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: get user failed: %w", err)
		}

		filI, err = f.searchStorage.GetFilterInterests(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: get interests failed: %w", err)
		}

		users, err = f.searchStorage.SearchUsers(ctx, fil, filI, page, limit)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: search failed: %w", err)
		}

		for _, u := range users {
			userInterest, err := f.userStorage.GetInterests(ctx, u.GetID())
			if err != nil {
				return fmt.Errorf("SearchUsersTx: get interests failed: %w", err)
			}

			u.Interest = userInterest
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("SearchUsersTx: user transaction failed: %w", err)
	}
	return users, nil
}
