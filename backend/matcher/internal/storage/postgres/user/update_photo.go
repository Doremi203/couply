package user

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) UpdatePhoto(ctx context.Context, photo user.Photo, userID uuid.UUID) error {
	const query = `
		UPDATE photos SET 
		    object_key = $3,
			mime_type = $4,
			uploaded_at = now()
		WHERE user_id = $1 AND order_number = $2
	`

	result, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx, query,
		userID, photo.OrderNumber, photo.ObjectKey, photo.MimeType,
	)
	if err != nil {
		return errors.WrapFail(err, "execute update photo query")
	}

	if result.RowsAffected() == 0 {
		return ErrPhotoNotFound
	}

	return nil
}

func (s *PgStorageUser) UpdatePhotoUploadedAt(ctx context.Context, orderNumbers []int32, userID uuid.UUID) error {
	const query = `
		UPDATE photos SET
			uploaded_at = now()
		WHERE user_id = $1 AND order_number = ANY($2)
	`

	result, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx, query,
		userID, orderNumbers,
	)
	if err != nil {
		return errors.WrapFail(err, "execute update photo uploaded at query")
	}

	if result.RowsAffected() == 0 {
		return ErrPhotoNotFound
	}

	return nil
}
