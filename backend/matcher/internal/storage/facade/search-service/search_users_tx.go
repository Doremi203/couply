package search_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeSearch) SearchUsersTx(ctx context.Context, userID int64, page, limit uint64) ([]*user.User, error) {
	var (
		fil   *search.Filter
		filI  *interest.Interest
		users []*user.User
		err   error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		fil, err = f.searchStorage.GetFilter(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "SearchUsersTx: get user failed")
		}

		filI, err = f.searchStorage.GetFilterInterests(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "SearchUsersTx: get interests failed")
		}

		users, err = f.searchStorage.SearchUsers(ctx, fil, filI, page, limit)
		if err != nil {
			return errors.Wrap(err, "SearchUsersTx: search failed")
		}

		for _, u := range users {
			userInterest, err := f.userStorage.GetInterests(ctx, u.GetID())
			if err != nil {
				return errors.Wrap(err, "SearchUsersTx: get interests failed")
			}

			u.Interest = userInterest
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "SearchUsersTx: user transaction failed")
	}
	return users, nil
}
