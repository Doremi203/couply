package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/user"
)

func (s *PgStorageUser) GetPhotos(ctx context.Context, userID int64) ([]*user.Photo, error) {
	photoSQL := `
		SELECT order_number, url, mime_type, uploaded_at, updated_at
		FROM Photos 
		WHERE user_id = $1
	`

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, photoSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("GetPhotos: %w", err)
	}
	defer rows.Close()

	photos := make([]*user.Photo, 0)
	for rows.Next() {
		var p user.Photo
		if err = rows.Scan(&p.OrderNumber, &p.URL, &p.MimeType, &p.UploadedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("GetPhotos scan: %w", err)
		}
		photos = append(photos, &p)
	}

	return photos, nil
}
