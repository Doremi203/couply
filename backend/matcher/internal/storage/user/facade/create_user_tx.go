package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) CreateUserTx(ctx context.Context, newUser *user.User) error {
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err = f.storage.CreateUser(ctxTx, newUser)
		if err != nil {
			return errors.Wrap(err, "storage.CreateUser")
		}

		for _, photo := range newUser.Photos {
			if err = f.storage.CreatePhoto(ctxTx, newUser.ID, photo); err != nil {
				return errors.Wrap(err, "storage.CreatePhoto")
			}
		}

		if err = f.storage.CreateInterests(ctxTx, newUser.ID, newUser.Interest); err != nil {
			return errors.Wrap(err, "storage.CreateInterests")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
