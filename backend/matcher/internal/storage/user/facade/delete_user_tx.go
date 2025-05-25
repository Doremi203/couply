package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	"github.com/google/uuid"
)

func (f *StorageFacadeUser) DeleteUserTx(ctx context.Context, userID uuid.UUID) error {
	err := f.storage.DeleteUser(ctx, userID)
	if err != nil {
		return errors.Wrap(err, "storage.DeleteUser")
	}

	return nil
}
