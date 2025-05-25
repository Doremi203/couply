package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	postgres2 "github.com/Doremi203/couply/backend/matcher/internal/storage/search/postgres"

	"github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeSearch) SearchUsersTx(ctx context.Context, userID uuid.UUID, offset, limit uint64) ([]*user.User, map[uuid.UUID]float64, error) {
	var (
		users     []*user.User
		distances map[uuid.UUID]float64
		err       error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		fil, err := f.searchStorage.GetFilter(ctxTx, postgres2.GetFilterOptions{
			UserID: userID,
		})
		if err != nil {
			return errors.Wrap(err, "searchStorage.GetFilter")
		}

		filI, err := f.searchStorage.GetFilterInterests(ctxTx, postgres2.GetFilterInterestsOptions{
			UserID: userID,
		})
		if err != nil {
			return errors.Wrap(err, "searchStorage.GetFilterInterests")
		}

		curUser, err := f.userStorage.GetUser(ctx, postgres.GetUserOptions{
			UserId: userID,
		})
		if err != nil {
			return errors.Wrap(err, "searchStorage.GetUser")
		}

		users, distances, err = f.searchStorage.SearchUsers(ctx, fil, filI, curUser.Latitude, curUser.Longitude, offset, limit)
		if err != nil {
			return errors.Wrap(err, "searchStorage.SearchUsers")
		}

		for _, u := range users {
			userInterest, err := f.userStorage.GetInterests(ctx, postgres.GetInterestsOptions{
				UserID: u.ID,
			})
			if err != nil {
				return errors.Wrap(err, "searchStorage.GetInterests")
			}

			u.Interest = userInterest
		}

		for _, u := range users {
			userPhotos, err := f.userStorage.GetPhotos(ctx, postgres.GetPhotosOptions{
				UserID: u.ID,
			})
			if err != nil {
				return errors.Wrap(err, "searchStorage.GetPhotos")
			}

			u.Photos = userPhotos
		}

		return nil
	})

	if err != nil {
		return nil, nil, errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return users, distances, nil
}
