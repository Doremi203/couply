package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageUser) DeletePhotos(ctx context.Context, userID uuid.UUID) error {
	query, args, err := buildDeletePhotosQuery(userID)
	if err != nil {
		return errors.Wrapf(err, "buildDeletePhotosQuery with %v", errors.Token("user_id", userID))
	}

	if err = executeDeletePhotosQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeDeletePhotosQuery with %v", errors.Token("user_id", userID))
	}

	return nil
}

func buildDeletePhotosQuery(userID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Delete(photosTableName).
		Where(sq.Eq{userIDColumnName: userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeDeletePhotosQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}
