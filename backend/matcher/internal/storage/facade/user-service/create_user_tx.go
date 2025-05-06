package user_service

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (f *StorageFacadeUser) CreateUserTx(ctx context.Context, newUser *user.User) (*user.User, error) {
	var createdUser *user.User
	var err error

	err = f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		createdUser, err = f.storage.AddUser(ctxTx, newUser)
		if err != nil {
			return fmt.Errorf("failed to add user: %w", err)
		}

		for _, photo := range createdUser.GetPhotos() {
			if err = f.storage.AddPhoto(ctxTx, photo, createdUser.GetID()); err != nil {
				return fmt.Errorf("failed to add photo: %w", err)
			}
		}

		if err = f.storage.AddInterests(ctxTx, createdUser.GetID(), createdUser.GetInterest()); err != nil {
			return fmt.Errorf("failed to add interests: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create user transaction: %w", err)
	}

	return createdUser, nil
}
