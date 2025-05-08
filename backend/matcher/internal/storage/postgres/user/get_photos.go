package user

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) GetPhotos(ctx context.Context, userID int64) ([]*user.Photo, error) {
	photoSQL := `
		SELECT order_number, url, mime_type, uploaded_at, updated_at
		FROM photos 
		WHERE user_id = $1
	`

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, photoSQL, userID)
	if err != nil {
		return nil, errors.Wrap(err, "GetPhotos")
	}
	defer rows.Close()

	photos := make([]*user.Photo, 0)
	for rows.Next() {
		var p user.Photo
		if err = rows.Scan(&p.OrderNumber, &p.URL, &p.MimeType, &p.UploadedAt, &p.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "GetPhotos scan")
		}
		photos = append(photos, &p)
	}

	return photos, nil
}
