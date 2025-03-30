package postgres

import (
	"context"
	"fmt"

	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (s *PgStorage) UpdatePhoto(ctx context.Context, photo *domain.Photo) error {
	photoSQL := `
        UPDATE Photos 
        SET url = $1, mime_type = $2, uploaded_at = $3
        WHERE user_id = $4 AND order_number = $5
    `

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		photoSQL,
		photo.URL,
		photo.MimeType,
		photo.UploadedAt,
		photo.UserID,
		photo.OrderNumber,
	)
	if err != nil {
		return fmt.Errorf("UpdatePhoto: %w", err)
	}

	return nil
}
