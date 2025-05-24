package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) CreateFilterTx(ctx context.Context, newFilter *search.Filter) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.searchStorage.CreateFilter(ctxTx, newFilter)
		if err != nil {
			return errors.Wrap(err, "searchStorage.CreateFilter")
		}

		if err = f.searchStorage.CreateFilterInterests(ctxTx, newFilter.UserID, newFilter.Interest); err != nil {
			return errors.Wrap(err, "searchStorage.CreateFilterInterests")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
