package user_service

import (
	"context"
)

func (f *StorageFacadeUser) DeleteUserTx(ctx context.Context, id int64) error {
	err := f.txManager.RunReadCommitted(ctx, func(ctx context.Context) error {
		err := f.storage.DeleteUser(ctx, id)
		return err
	})

	return err
}
