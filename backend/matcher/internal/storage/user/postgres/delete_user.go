package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageUser) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	query, args, err := buildDeleteUserQuery(userID)
	if err != nil {
		return errors.Wrapf(err, "buildDeleteUserQuery with %v", errors.Token("user_id", userID))
	}

	if err = executeDeleteUserQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeDeleteUserQuery with %v", errors.Token("user_id", userID))
	}

	return nil
}

func buildDeleteUserQuery(userID uuid.UUID) (string, []any, error) {
	query, args, err := sq.Delete(usersTableName).
		Where(sq.Eq{idColumnName: userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeDeleteUserQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}
