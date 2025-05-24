package facade

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/google/uuid"
)

func (f *StorageFacadeMatching) DeleteMatchTx(ctx context.Context, userID, targetUserID uuid.UUID) error {
	if err := f.storage.DeleteMatch(ctx, userID, targetUserID); err != nil {
		return errors.Wrap(err, "storage.DeleteMatch")
	}

	return nil
}
