package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

func (f *StorageFacadeUser) GetUsersTx(ctx context.Context, userIDs []uuid.UUID) ([]*user.User, error) {
	var (
		users []*user.User
		err   error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		users, err = f.storage.GetUsers(ctxTx, postgres.GetUsersOptions{
			UserIDs: userIDs,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetUsers")
		}

		photosMap, err := f.storage.GetMultiplePhotos(ctxTx, postgres.GetMultiplePhotosOptions{
			UserIDs: userIDs,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetMultiplePhotos")
		}

		interestsMap, err := f.storage.GetMultipleInterests(ctxTx, postgres.GetMultipleInterestsOptions{
			UserIDs: userIDs,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetMultipleInterests")
		}

		for _, u := range users {
			if photos, exists := photosMap[u.ID]; exists {
				u.Photos = photos
			}

			if interest, exists := interestsMap[u.ID]; exists {
				u.Interest = interest
			}
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return users, nil
}
