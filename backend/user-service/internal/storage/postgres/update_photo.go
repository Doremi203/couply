package postgres

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (s *PgStorage) UpdatePhoto(ctx context.Context, userID int64, photo *domain.Photo) error {
	const op = "UpdatePhoto"

	photoSQL := `
        UPDATE Photos 
        SET url = $1, mime_type = $2, uploaded_at = $3
        WHERE id = $4 AND user_id = $5
    `

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		photoSQL,
		photo.URL,
		photo.MimeType,
		photo.UploadedAt,
		photo.ID,
		userID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
