package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) AddPhoto(ctx context.Context, photo *user.Photo, userID int64) error {
	photoSQL := `
		INSERT INTO photos (user_id, order_number, url, mime_type, uploaded_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		photoSQL,
		userID,
		photo.OrderNumber,
		photo.URL,
		photo.MimeType,
		photo.UploadedAt,
		photo.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("AddPhoto: %w", err)
	}

	return nil
}
