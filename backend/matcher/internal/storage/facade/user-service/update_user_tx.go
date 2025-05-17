package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) UpdateUserTx(ctx context.Context, user *user.User) (*user.User, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if _, err := f.storage.UpdateUser(ctxTx, user); err != nil {
			return errors.WrapFail(err, "update user")
		}

		for _, photo := range user.GetPhotos() {
			if err := f.storage.UpdatePhoto(ctxTx, photo, user.GetID()); err != nil {
				return errors.WrapFail(err, "update photo")
			}
		}

		if err := f.storage.DeleteInterests(ctxTx, user.GetID()); err != nil {
			return errors.WrapFail(err, "delete old interests")
		}

		if err := f.storage.AddInterests(ctxTx, user.GetID(), user.GetInterest()); err != nil {
			return errors.WrapFail(err, "add new interests")
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "UpdateUserTx failed")
	}
	return user, nil
}
