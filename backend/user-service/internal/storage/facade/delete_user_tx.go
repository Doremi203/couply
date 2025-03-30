package facade

import "context"

func (f *StorageFacade) DeleteUserTx(ctx context.Context, id int64) error {
	err := f.txManager.RunReadCommitted(ctx, func(ctx context.Context) error {
		err := f.storage.DeleteUser(ctx, id)
		return err
	})

	return err
}
