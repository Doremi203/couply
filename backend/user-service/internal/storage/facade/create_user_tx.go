package facade

import (
	"context"
	"fmt"

	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (f *StorageFacade) CreateUserTx(ctx context.Context, user *domain.User) (*domain.User, error) {
	var createdUser *domain.User
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		createdUser, err = f.storage.AddUser(ctxTx, user)
		if err != nil {
			return fmt.Errorf("failed to add user: %w", err)
		}

		for _, photo := range user.Photos {
			if err = f.storage.AddPhoto(ctxTx, photo); err != nil {
				return fmt.Errorf("failed to add photo: %w", err)
			}
		}

		if err = f.storage.AddInterests(ctxTx, user.ID, user.Interest); err != nil {
			return fmt.Errorf("failed to add interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create user transaction: %w", err)
	}

	return createdUser, nil
}
