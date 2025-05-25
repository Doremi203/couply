package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) UpdatePhoto(ctx context.Context, photo user.Photo) error {
	query, args, err := buildUpdatePhotoQuery(photo)
	if err != nil {
		return errors.Wrapf(err, "buildUpdatePhotoQuery with %v", errors.Token("user_id", photo.UserID))
	}

	result, err := executeUpdatePhotoQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeUpdatePhotoQuery with %v", errors.Token("user_id", photo.UserID))
	}

	if result.RowsAffected() == 0 {
		return user.ErrPhotoNotFound
	}

	return nil
}

func buildUpdatePhotoQuery(photo user.Photo) (string, []any, error) {
	query, args, err := sq.Update(photosTableName).
		Set(objectKeyColumnName, photo.ObjectKey).
		Set(mimeTypeColumnName, photo.MimeType).
		Set(uploadedAtColumnName, photo.UploadedAt).
		Where(sq.And{
			sq.Eq{userIDColumnName: photo.UserID},
			sq.Eq{orderNumberColumnName: photo.OrderNumber},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeUpdatePhotoQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}
