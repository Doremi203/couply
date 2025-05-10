package user

import (
	"context"
	"fmt"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	userpkg "github.com/Doremi203/couply/backend/auth/pkg/user"
	"github.com/Doremi203/couply/backend/common/libs/slices"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type photoEntity struct {
	UserID      uuid.UUID  `db:"user_id"`
	OrderNumber int32      `db:"order_number"`
	ObjectKey   string     `db:"object_key"`
	MimeType    string     `db:"mime_type"`
	UploadedAt  *time.Time `db:"uploaded_at"`
}

func (s *PgStorageUser) GetPhotos(ctx context.Context, userID uuid.UUID) ([]user.Photo, error) {
	query, args, err := sq.Select(
		"user_id", "order_number", "object_key", "mime_type", "uploaded_at",
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

	photos, err := pgx.CollectRows(rows, pgx.RowToStructByName[photoEntity])
	if err != nil {
		return nil, errors.WrapFail(err, "collect photos rows")
	}

	return slices.Map(photos, func(from photoEntity) user.Photo {
		return user.Photo{
			UserID:      userpkg.ID(from.UserID),
			OrderNumber: from.OrderNumber,
			ObjectKey:   from.ObjectKey,
			MimeType:    from.MimeType,
			UploadedAt:  from.UploadedAt,
		}
	}), nil
}
