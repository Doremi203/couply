package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"
)

func (s *PgStorageSearch) DeleteFilterInterests(ctx context.Context, userID uuid.UUID) error {
	query, args, err := buildDeleteFilterInterestsQuery(userID)
	if err != nil {
		return errors.Wrapf(err, "buildDeleteFilterInterestsQuery with %v", errors.Token("user_id", userID))
	}

	if err = executeDeleteFilterInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeDeleteFilterInterestsQuery with %v", errors.Token("user_id", userID))
	}

	return nil
}

func buildDeleteFilterInterestsQuery(userID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Delete(filterInterestsTableName).
		Where(sq.Eq{userIdColumnName: userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeDeleteFilterInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}
