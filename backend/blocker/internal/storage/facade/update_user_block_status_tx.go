package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/google/uuid"
)

func (f *StorageFacadeBlocker) UpdateUserBlockStatusTx(ctx context.Context, blockID uuid.UUID, status blocker.BlockStatus) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := f.storage.UpdateUserBlockStatus(ctx, blockID, status)
		return err
	})

	return err
}
