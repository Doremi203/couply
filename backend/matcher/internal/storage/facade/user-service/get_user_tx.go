package user_service

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) GetUserTx(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	var (
		u      *user.User
		photos []user.Photo
		i      *interest.Interest
		err    error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		u, err = f.storage.GetUser(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("GetUserTx: get user failed: %w", err)
		}

		photos, err = f.storage.GetPhotos(ctxTx, userID)
		if err != nil {
			return errors.Wrap(err, "GetUserTx: get photos failed")
		}

		i, err = f.storage.GetInterests(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("GetUserTx: get interests failed: %w", err)
		}

		u.Photos = photos
		u.Interest = i

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("GetUserTx: user transaction failed: %w", err)
	}
	return u, nil
}
