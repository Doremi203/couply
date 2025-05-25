package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageUser) DeleteInterests(ctx context.Context, userID uuid.UUID) error {
	query, args, err := buildDeleteInterestsQuery(userID)
	if err != nil {
		return errors.Wrapf(err, "buildDeleteInterestsQuery with %v", errors.Token("user_id", userID))
	}

	result, err := executeDeleteInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeDeleteInterestsQuery with %v", errors.Token("user_id", userID))
	}

	if result.RowsAffected() == 0 {
		return interest.ErrInterestsNotFound
	}

	return nil
}

func buildDeleteInterestsQuery(userID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Delete(interestsTableName).
		Where(sq.Eq{userIDColumnName: userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeDeleteInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}
