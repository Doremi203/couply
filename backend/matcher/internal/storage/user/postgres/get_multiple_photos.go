package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type GetMultiplePhotosOptions struct {
	UserIDs []uuid.UUID
}

func (s *PgStorageUser) GetMultiplePhotos(ctx context.Context, opts GetMultiplePhotosOptions) (map[uuid.UUID][]user.Photo, error) {
	if len(opts.UserIDs) == 0 {
		return map[uuid.UUID][]user.Photo{}, nil
	}

	query, args, err := buildGetMultiplePhotosQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetMultiplePhotosQuery with %v", errors.Token("options", opts))
	}

	photosFromDB, err := executeGetMultiplePhotosQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetMultiplePhotosQuery with %v", errors.Token("options", opts))
	}

	photosMap := make(map[uuid.UUID][]user.Photo)
	for _, p := range photosFromDB {
		userPhoto := user.Photo{
			UserID:      p.UserID,
			OrderNumber: p.OrderNumber,
			ObjectKey:   p.ObjectKey,
			MimeType:    p.MimeType,
			UploadedAt:  p.UploadedAt,
		}
		photosMap[userPhoto.UserID] = append(photosMap[userPhoto.UserID], userPhoto)
	}

	return photosMap, nil
}

func buildGetMultiplePhotosQuery(opts GetMultiplePhotosOptions) (string, []any, error) {
	query, args, err := sq.Select(photosColumns...).
		From(photosTableName).
		Where(sq.Eq{userIDColumnName: opts.UserIDs}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetMultiplePhotosQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]dbPhoto, error) {
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
