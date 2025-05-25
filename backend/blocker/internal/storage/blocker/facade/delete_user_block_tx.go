package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage/blocker/postgres"

	"github.com/google/uuid"
)

func (f *StorageFacadeBlocker) DeleteUserBlockTx(ctx context.Context, userID uuid.UUID) error {
	if err := f.storage.DeleteUserBlock(ctx, postgres.DeleteUserBlockOptions{
		UserID: userID,
	}); err != nil {
		return errors.Wrap(err, "storage.DeleteUserBlock")
	}

	return nil
}
