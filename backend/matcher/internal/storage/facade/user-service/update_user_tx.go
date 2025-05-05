package user_service

import (
	"context"
	"fmt"
	"time"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) UpdateUserTx(ctx context.Context, user *user.User) (*user.User, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if _, err := f.storage.UpdateUser(ctxTx, user); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		for _, photo := range user.GetPhotos() {
			photo.UpdatedAt = time.Now()
			if err := f.storage.UpdatePhoto(ctxTx, photo, user.GetID()); err != nil {
				return fmt.Errorf("failed to update photo: %w", err)
			}
		}

		if err := f.storage.DeleteInterests(ctxTx, user.GetID()); err != nil {
			return fmt.Errorf("failed to delete old interests: %w", err)
		}

		if err := f.storage.AddInterests(ctxTx, user.GetID(), user.GetInterest()); err != nil {
			return fmt.Errorf("failed to add new interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("UpdateUserTx failed: %w", err)
	}
	return user, nil
}
