package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/blocker/postgres"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/google/uuid"
)

func (f *StorageFacadeBlocker) UpdateUserBlockStatusTx(ctx context.Context, blockID uuid.UUID, status blocker.BlockStatus) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		block, err := f.storage.GetUserBlock(ctxTx, postgres.GetUserBlockOptions{
			BlockID:   blockID,
			ForUpdate: true,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetUserBlock")
		}

		block.Status = status

		err = f.storage.UpdateUserBlock(ctxTx, block)
		if err != nil {
			return errors.Wrap(err, "storage.UpdateUserBlock")
		}

		return nil
	})

	return err
}
