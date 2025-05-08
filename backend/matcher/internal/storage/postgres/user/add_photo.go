package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) AddPhoto(ctx context.Context, photo *user.Photo, userID uuid.UUID) error {
	query, args, err := sq.Insert("photos").
		Columns("user_id", "order_number", "url", "mime_type", "uploaded_at", "updated_at").
		Values(
			userID,
			photo.GetOrderNumber(),
			photo.GetURL(),
			photo.GetMimeType(),
			photo.GetUploadedAt(),
			photo.GetUpdatedAt(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicatePhoto
		}
		return fmt.Errorf("failed to insert photo: %w", err)
	}

	return nil
}
