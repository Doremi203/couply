package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageUser) GetPhotosForUsers(ctx context.Context, userIDs []uuid.UUID) (map[uuid.UUID][]user.Photo, error) {
	if len(userIDs) == 0 {
		return map[uuid.UUID][]user.Photo{}, nil
	}

	query, args, err := sq.Select(
		"user_id", "order_number", "object_key", "mime_type", "uploaded_at",
	).
		From("photos").
		Where(sq.Eq{"user_id": userIDs}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build photos query: %w", err)
	}

	rows, err := s.txManager.GetQueryEngine(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute photos query: %w", err)
	}
	defer rows.Close()

	photosMap := make(map[uuid.UUID][]user.Photo)
	for rows.Next() {
		var (
			userID uuid.UUID
			p      user.Photo
		)
		err := rows.Scan(
			&userID,
			&p.OrderNumber,
			&p.ObjectKey,
			&p.MimeType,
			&p.UploadedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan photo row: %w", err)
		}
		photosMap[userID] = append(photosMap[userID], p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("photos rows iteration error: %w", err)
	}

	return photosMap, nil
}
