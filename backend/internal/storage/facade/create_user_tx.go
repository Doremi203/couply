package facade

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (f *StorageFacade) CreateUserTx(ctx context.Context, user domain.User) error {
	err := f.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		userID, err := f.storage.AddUser(ctxTx, user)
		if err != nil {
			return fmt.Errorf("failed to add user: %w", err)
		}

		for _, photo := range user.Photos {
			if err := f.storage.AddPhoto(ctxTx, userID, *photo); err != nil {
				return fmt.Errorf("failed to add photo: %w", err)
			}
		}

		if err := f.storage.AddInterests(ctxTx, userID, user.Interest); err != nil {
			return fmt.Errorf("failed to add interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to create user transaction: %w", err)
	}

	return nil
}
