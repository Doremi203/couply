package search_service

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeSearch) SearchUsersTx(ctx context.Context, userID int64, page, limit int32) ([]*user.User, error) {
	var (
		fil   *search.Filter
		i     *interest.Interest
		users []*user.User
		err   error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		fil, err = f.storage.GetFilter(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: get user failed: %w", err)
		}

		i, err = f.storage.GetFilterInterests(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: get interests failed: %w", err)
		}

		users, err = f.storage.SearchUsers(ctx, fil, i, page, limit)
		if err != nil {
			return fmt.Errorf("SearchUsersTx: search failed: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("SearchUsersTx: user transaction failed: %w", err)
	}
	return users, nil
}
