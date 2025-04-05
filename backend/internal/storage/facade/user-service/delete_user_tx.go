package user_service

import (
	"context"
)

func (f *StorageFacadeUser) DeleteUserTx(ctx context.Context, id int64) error {
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
