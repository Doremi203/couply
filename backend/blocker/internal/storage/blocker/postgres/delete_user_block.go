package postgres

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type DeleteUserBlockOptions struct {
	UserID uuid.UUID
}

func (s *PgStorageBlocker) DeleteUserBlock(ctx context.Context, opts DeleteUserBlockOptions) error {
	query, args, err := buildDeleteUserBlockQuery(opts)
	if err != nil {
		return errors.Wrapf(err, "buildDeleteUserBlockQuery with %v", errors.Token("options", opts))
	}

	if err = executeDeleteUserBlockQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeDeleteUserBlockQuery with %v", errors.Token("options", opts))
	}

	return nil
}

func buildDeleteUserBlockQuery(opts DeleteUserBlockOptions) (string, []any, error) {
	query, args, err := sq.Delete(userBlocksTableName).
		Where(sq.Eq{blockedIdColumn: opts.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeDeleteUserBlockQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}
