package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/blocker/postgres"

	"github.com/google/uuid"
)

func (f *StorageFacadeBlocker) DeleteUserBlockTx(ctx context.Context, userID uuid.UUID) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := f.storage.DeleteUserBlock(ctx, postgres.DeleteUserBlockOptions{
			UserID: userID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.DeleteUserBlock")
		}

		return nil
	})

	return err
}
