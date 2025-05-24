package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) UpdateFilterTx(ctx context.Context, filter *search.Filter) (*search.Filter, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if err := f.searchStorage.UpdateFilter(ctxTx, filter); err != nil {
			return errors.WrapFail(err, "update filter")
		}

		if err := f.searchStorage.DeleteFilterInterests(ctxTx, filter.UserID); err != nil {
			return errors.WrapFail(err, "delete old filter interests")
		}

		if err := f.searchStorage.AddFilterInterests(ctxTx, filter.UserID, filter.Interest); err != nil {
			return errors.WrapFail(err, "add new filter interests")
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "UpdateFilterTx failed")
	}
	return filter, nil
}
