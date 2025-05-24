package postgres

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/common/libs/slices"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type dbPhoto struct {
	UserID      uuid.UUID  `db:"user_id"`
	OrderNumber int32      `db:"order_number"`
	ObjectKey   string     `db:"object_key"`
	MimeType    string     `db:"mime_type"`
	UploadedAt  *time.Time `db:"uploaded_at"`
}

type GetPhotosOptions struct {
	UserID       uuid.UUID
	OrderNumbers []int32
	ForUpdate    bool
}

func (s *PgStorageUser) GetPhotos(ctx context.Context, opts GetPhotosOptions) ([]user.Photo, error) {
	query, args, err := buildGetPhotosQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetPhotosQuery with %v", errors.Token("options", opts))
	}

	photos, err := executeGetPhotosQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetPhotosQuery with %v", errors.Token("options", opts))
	}

	return slices.Map(photos, func(from dbPhoto) user.Photo {
		return user.Photo{
			UserID:      from.UserID,
			OrderNumber: from.OrderNumber,
			ObjectKey:   from.ObjectKey,
			MimeType:    from.MimeType,
			UploadedAt:  from.UploadedAt,
		}
	}), nil
}

func buildGetPhotosQuery(opts GetPhotosOptions) (string, []any, error) {
	sb := sq.Select(photosColumns...).
		From(photosTableName).
		Where(sq.Eq{userIDColumnName: opts.UserID})

	if len(opts.OrderNumbers) != 0 {
		sb = sb.Where(userIDColumnName, opts.UserID).
			Where(orderNumberColumnName, opts.OrderNumbers)
	} else {
		sb = sb.Where(userIDColumnName, opts.UserID)
	}

	if opts.ForUpdate {
		sb = sb.Suffix("FOR UPDATE")
	}

	return sb.PlaceholderFormat(sq.Dollar).ToSql()
}

func executeGetPhotosQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]dbPhoto, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	photos, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbPhoto])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(user.ErrPhotosNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return photos, nil
}
