package facade

import (
	"context"

	"github.com/google/uuid"
)

func (f *StorageFacadeBlocker) DeleteUserBlockTx(ctx context.Context, userID uuid.UUID) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := f.storage.DeleteUserBlock(ctx, userID)
		return err
	})

	return err
}
