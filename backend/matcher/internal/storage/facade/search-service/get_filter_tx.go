package search_service

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) GetFilterTx(ctx context.Context, userID int64) (*search.Filter, error) {
	var (
		fil *search.Filter
		i   *interest.Interest
		err error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		fil, err = f.storage.GetFilter(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("GetFilterTx: get user failed: %w", err)
		}

		i, err = f.storage.GetFilterInterests(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("GetFilterTx: get interests failed: %w", err)
		}

		fil.Interest = i

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("GetFilterTx: user transaction failed: %w", err)
	}
	return fil, nil
}
