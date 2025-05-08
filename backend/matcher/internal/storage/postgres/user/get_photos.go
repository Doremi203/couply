package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) GetPhotos(ctx context.Context, userID uuid.UUID) ([]*user.Photo, error) {
	query, args, err := sq.Select(
		"order_number", "url", "mime_type", "uploaded_at", "updated_at",
	).
		From("photos").
		Where(sq.Eq{"user_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	photos := make([]*user.Photo, 0)
	for rows.Next() {
		var p user.Photo
		if err := rows.Scan(
			&p.OrderNumber,
			&p.URL,
			&p.MimeType,
			&p.UploadedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		photos = append(photos, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return photos, nil
}
