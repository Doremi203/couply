package postgres

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (s *PgStorage) AddPhoto(ctx context.Context, userID int64, photo domain.Photo) error {
	const op = "AddPhotos"

	photoSQL := `
		INSERT INTO Photos (user_id, url, mime_type, uploaded_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, photoSQL, userID, photo.URL, photo.MimeType, photo.UploadedAt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
