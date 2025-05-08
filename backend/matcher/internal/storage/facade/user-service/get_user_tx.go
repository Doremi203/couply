package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) GetUserTx(ctx context.Context, userID int64) (*user.User, error) {
	var (
		u   *user.User
		p   []*user.Photo
		i   *interest.Interest
		err error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		u, err = f.storage.GetUser(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "GetUserTx: get user failed")
		}

		p, err = f.storage.GetPhotos(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "GetUserTx: get photos failed")
		}

		i, err = f.storage.GetInterests(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "GetUserTx: get interests failed")
		}

		u.Photos = p
		u.Interest = i

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "GetUserTx: user transaction failed")
	}
	return u, nil
}
