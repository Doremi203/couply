package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) CreateFilterTx(ctx context.Context, newFilter *search.Filter) (*search.Filter, error) {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.searchStorage.AddFilter(ctxTx, newFilter)
		if err != nil {
			return errors.WrapFail(err, "add filter")
		}

		if err = f.searchStorage.AddFilterInterests(ctxTx, newFilter.GetUserID(), newFilter.GetInterest()); err != nil {
			return errors.WrapFail(err, "add filter interests")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "create filter transaction")
	}

	return newFilter, nil
}
