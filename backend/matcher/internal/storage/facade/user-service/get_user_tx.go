package user_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) GetUserTx(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	var (
		u   *user.User
		p   []*user.Photo
		i   *interest.Interest
		err error
	)

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		u, err = f.storage.GetUser(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("GetUserTx: get user failed: %w", err)
		}

		p, err = f.storage.GetPhotos(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("GetUserTx: get photos failed: %w", err)
		}

		i, err = f.storage.GetInterests(ctxTx, userID)
		if err != nil {
			return fmt.Errorf("GetUserTx: get interests failed: %w", err)
		}

		u.Photos = p
		u.Interest = i

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("GetUserTx: user transaction failed: %w", err)
	}
	return u, nil
}
