package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) CreateUserTx(ctx context.Context, newUser *user.User) (*user.User, error) {
	var createdUser *user.User
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		createdUser, err = f.storage.AddUser(ctxTx, newUser)
		if err != nil {
			return errors.WrapFail(err, "add user")
		}

		for _, photo := range createdUser.Photos {
			if err = f.storage.AddPhoto(ctxTx, photo, createdUser.ID); err != nil {
				return errors.WrapFail(err, "add photo")
			}
		}

		if err = f.storage.AddInterests(ctxTx, createdUser.ID, createdUser.Interest); err != nil {
			return errors.WrapFail(err, "add interests")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WrapFail(err, "create user transaction")
	}

	return createdUser, nil
}
