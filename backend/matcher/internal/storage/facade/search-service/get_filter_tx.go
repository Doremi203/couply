package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
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
		fil, err = f.searchStorage.GetFilter(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "GetFilterTx: get user failed")
		}

		i, err = f.searchStorage.GetFilterInterests(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "GetFilterTx: get interests failed")
		}

		fil.Interest = i

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "GetFilterTx: user transaction failed")
	}
	return fil, nil
}
