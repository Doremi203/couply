package search_service

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) CreateFilterTx(ctx context.Context, newFilter *search.Filter) (*search.Filter, error) {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.searchStorage.AddFilter(ctxTx, newFilter)
		if err != nil {
			return fmt.Errorf("failed to add filter: %w", err)
		}

		if err = f.searchStorage.AddFilterInterests(ctxTx, newFilter.GetUserID(), newFilter.GetInterest()); err != nil {
			return fmt.Errorf("failed to add filter interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create filter transaction: %w", err)
	}

	return newFilter, nil
}
