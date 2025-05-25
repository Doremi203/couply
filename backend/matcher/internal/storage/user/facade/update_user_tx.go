package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) UpdateUserTx(ctx context.Context, user *user.User) (*user.User, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if err := f.storage.UpdateUser(ctxTx, user); err != nil {
			return errors.Wrap(err, "storage.UpdateUser")
		}

		for _, photo := range user.Photos {
			if err := f.storage.UpdatePhoto(ctxTx, photo); err != nil {
				return errors.Wrap(err, "storage.UpdatePhoto")
			}
		}

		if err := f.storage.DeleteInterests(ctxTx, user.ID); err != nil {
			return errors.Wrap(err, "storage.DeleteInterests")
		}

		if err := f.storage.CreateInterests(ctxTx, user.ID, user.Interest); err != nil {
			return errors.Wrap(err, "storage.CreateInterests")
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return user, nil
}
