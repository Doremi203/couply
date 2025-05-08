package user_service

import (
	"context"

	"github.com/google/uuid"
)

func (f *StorageFacadeUser) DeleteUserTx(ctx context.Context, id uuid.UUID) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := f.storage.DeleteUser(ctx, id)
		if err != nil {
			return err
		}

		err = f.storage.DeletePhotos(ctx, id)
		if err != nil {
			return err
		}

		err = f.storage.DeleteInterests(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return err
}
