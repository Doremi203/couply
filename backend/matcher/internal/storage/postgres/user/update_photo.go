package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) UpdatePhoto(ctx context.Context, photo *user.Photo, userID int64) error {
	photoSQL := `
        UPDATE photos 
        SET url = $1, mime_type = $2, updated_at = $3
        WHERE user_id = $4 AND order_number = $5
    `

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		photoSQL,
		photo.URL,
		photo.MimeType,
		photo.UpdatedAt,
		userID,
		photo.OrderNumber,
	)
	if err != nil {
		return fmt.Errorf("UpdatePhoto: %w", err)
	}

	return nil
}
