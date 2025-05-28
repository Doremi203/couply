package facade

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"
	"github.com/google/uuid"
)

func (f *StorageFacadeUser) UpdateVerificationStatusTx(
	ctx context.Context,
	userID uuid.UUID,
	isVerified bool,
) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		u, err := f.storage.GetUser(ctx, postgres.GetUserOptions{
			UserId: userID,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetUser")
		}

		u.IsVerified = isVerified

		if err := f.storage.UpdateUser(ctxTx, u); err != nil {
			return errors.Wrap(err, "storage.UpdateUser")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
