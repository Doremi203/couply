package matching_service

import (
	"context"

	"github.com/google/uuid"
)

func (f *StorageFacadeMatching) DeleteMatchTx(ctx context.Context, userID, targetUserID uuid.UUID) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := f.storage.DeleteMatch(ctx, userID, targetUserID)
		return err
	})

	return err
}
