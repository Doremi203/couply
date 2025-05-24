package facade

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/matcher/internal/storage/user/postgres"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

func (f *StorageFacadeUser) UpdatePhotosUploadedAtTx(ctx context.Context, orderNumbers []int32, userID uuid.UUID) error {
	err := f.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		photos, err := f.storage.GetPhotos(ctxTx, postgres.GetPhotosOptions{
			UserID:       userID,
			OrderNumbers: orderNumbers,
			ForUpdate:    true,
		})
		if err != nil {
			return errors.Wrap(err, "storage.GetPhotos")
		}

		now := time.Now()
		for _, photo := range photos {
			photo.UploadedAt = &now
			if err = f.storage.UpdatePhoto(ctxTx, photo); err != nil {
				return errors.WrapFail(err, "update user")
			}
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "txManager.RunRepeatableRead")
	}

	return nil
}
