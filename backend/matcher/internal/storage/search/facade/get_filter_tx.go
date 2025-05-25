package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/Doremi203/couply/backend/matcher/internal/storage/search/postgres"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (f *StorageFacadeSearch) GetFilterTx(ctx context.Context, userID uuid.UUID) (*search.Filter, error) {
	var (
		fil *search.Filter
		i   *interest.Interest
		err error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		fil, err = f.searchStorage.GetFilter(ctxTx, postgres.GetFilterOptions{
			UserID: userID,
		})
		if err != nil {
			return errors.Wrap(err, "searchStorage.GetFilter")
		}

		i, err = f.searchStorage.GetFilterInterests(ctxTx, postgres.GetFilterInterestsOptions{
			UserID: userID,
		})
		if err != nil {
			return errors.Wrap(err, "searchStorage.GetFilterInterests")
		}

		fil.Interest = i

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return fil, nil
}
