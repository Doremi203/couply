package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) CreatePhoto(ctx context.Context, userID uuid.UUID, photo user.Photo) error {
	query, args, err := buildCreatePhotoQuery(userID, photo)
	if err != nil {
		return errors.Wrapf(err, "buildCreatePhotoQuery with %v", errors.Token("user_id", userID))
	}

	if err = executeCreatePhotoQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreatePhotoQuery with %v", errors.Token("user_id", userID))
	}

	return nil
}

func buildCreatePhotoQuery(userID uuid.UUID, photo user.Photo) (string, []any, error) {
	query, args, err := sq.Insert(photosTableName).
		Columns(photosColumns...).
		Values(userID, photo.OrderNumber, photo.ObjectKey, photo.MimeType, photo.UploadedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreatePhotoQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if postgres.IsForeignKeyViolationError(err) {
			return user.ErrUserDoesntExist
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
