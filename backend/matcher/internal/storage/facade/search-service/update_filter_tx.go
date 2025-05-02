package search_service

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) UpdateFilterTx(ctx context.Context, filter *search.Filter) (*search.Filter, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if err := f.storage.UpdateFilter(ctxTx, filter); err != nil {
			return fmt.Errorf("failed to update filter: %w", err)
		}

		if err := f.storage.DeleteFilterInterests(ctxTx, filter.UserID); err != nil {
			return fmt.Errorf("failed to delete old filter interests: %w", err)
		}

		if err := f.storage.AddFilterInterests(ctxTx, filter.UserID, filter.Interest); err != nil {
			return fmt.Errorf("failed to add new filter interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("UpdateFilterTx failed: %w", err)
	}
	return filter, nil
}
