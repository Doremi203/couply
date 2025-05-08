package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) UpdatePhoto(ctx context.Context, photo *user.Photo, userID uuid.UUID) error {
	query, args, err := sq.Update("photos").
		Set("url", photo.GetURL()).
		Set("mime_type", photo.GetMimeType()).
		Set("updated_at", photo.GetUpdatedAt()).
		Where(sq.Eq{
			"user_id":      userID,
			"order_number": photo.GetOrderNumber(),
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	result, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrPhotoNotFound
	}

	return nil
}
