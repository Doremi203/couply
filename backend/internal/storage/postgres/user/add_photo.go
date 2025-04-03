package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/user"
)

func (s *PgStorageUser) AddPhoto(ctx context.Context, photo *user.Photo) error {
	photoSQL := `
		INSERT INTO Photos (user_id, order_number, url, mime_type, uploaded_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		photoSQL,
		photo.UserID,
		photo.OrderNumber,
		photo.URL,
		photo.MimeType,
		photo.UploadedAt,
	)
	if err != nil {
		return fmt.Errorf("AddPhoto: %w", err)
	}

	return nil
}
