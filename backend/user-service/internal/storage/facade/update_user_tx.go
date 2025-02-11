package facade

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (f *StorageFacade) UpdateUserTx(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		_, err := f.storage.UpdateUser(ctxTx, user)
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		if err = f.storage.DeletePhotos(ctxTx, user.ID); err != nil {
			return fmt.Errorf("failed to delete old photos: %w", err)
		}
		for _, photo := range user.Photos {
			if err = f.storage.AddPhoto(ctxTx, user.ID, photo); err != nil {
				return fmt.Errorf("failed to add new photo: %w", err)
			}
		}

		if err = f.storage.UpdateInterests(ctxTx, user.ID, user.Interest); err != nil {
			return fmt.Errorf("failed to update interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update user transaction: %w", err)
	}

	return user, nil
}
