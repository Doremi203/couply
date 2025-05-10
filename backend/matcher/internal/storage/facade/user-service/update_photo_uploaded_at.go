package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

func (f *StorageFacadeUser) UpdatePhotosUploadedAtTx(
	ctx context.Context,
	orderNumbers []int32,
	userID uuid.UUID,
) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		if err := f.storage.UpdatePhotoUploadedAt(ctxTx, orderNumbers, userID); err != nil {
			return errors.WrapFail(err, "update user")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "UpdatePhotosUploadedAtTx failed")
	}

	return nil
}
