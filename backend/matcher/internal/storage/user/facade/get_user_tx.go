package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) GetUserTx(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	var (
		domainUser *user.User
		err        error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		domainUser, err = f.storage.GetUser(ctxTx, postgres.GetUserOptions{
			UserId: userID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetUser")
		}

		p, err := f.storage.GetPhotos(ctxTx, postgres.GetPhotosOptions{
			UserID: userID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetPhotos")
		}

		i, err := f.storage.GetInterests(ctxTx, postgres.GetInterestsOptions{
			UserID: userID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetInterests")
		}

		domainUser.Photos = p
		domainUser.Interest = i

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return domainUser, nil
}
