package search_service

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) UpdateFilterTx(ctx context.Context, filter *search.Filter) (*search.Filter, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if err := f.searchStorage.UpdateFilter(ctxTx, filter); err != nil {
			return fmt.Errorf("failed to update filter: %w", err)
		}

		if err := f.searchStorage.DeleteFilterInterests(ctxTx, filter.GetUserID()); err != nil {
			return fmt.Errorf("failed to delete old filter interests: %w", err)
		}

		if err := f.searchStorage.AddFilterInterests(ctxTx, filter.GetUserID(), filter.GetInterest()); err != nil {
			return fmt.Errorf("failed to add new filter interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("UpdateFilterTx failed: %w", err)
	}
	return filter, nil
}
