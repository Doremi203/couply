package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) UpdateFilterTx(ctx context.Context, filter *search.Filter) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if err := f.searchStorage.UpdateFilter(ctxTx, filter); err != nil {
			return errors.Wrap(err, "searchStorage.UpdateFilter")
		}

		if err := f.searchStorage.DeleteFilterInterests(ctxTx, filter.UserID); err != nil {
			return errors.Wrap(err, "searchStorage.DeleteFilterInterests")
		}

		if err := f.searchStorage.CreateFilterInterests(ctxTx, filter.UserID, filter.Interest); err != nil {
			return errors.Wrap(err, "searchStorage.CreateFilterInterests")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
