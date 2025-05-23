package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageBlocker) UpdateUserBlock(ctx context.Context, block *blocker.UserBlock) error {
	query, args, err := buildUpdateUserBlockQuery(block)
	if err != nil {
		return errors.Wrapf(err, "buildUpdateUserBlockQuery with %v", errors.Token("block", block))
	}

	result, err := executeUpdateUserBlockQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeUpdateUserBlockQuery with %v", errors.Token("block", block))
	}

	if err = verifyUpdateResult(result); err != nil {
		return errors.Wrapf(err, "verifyUpdateResult with %v", errors.Token("block", block))
	}

	return nil
}

func buildUpdateUserBlockQuery(block *blocker.UserBlock) (string, []any, error) {
	query, args, err := sq.Update(userBlocksTableName).
		Set(statusColumn, block.Status).
		Where(sq.Eq{idColumn: block.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeUpdateUserBlockQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}

func verifyUpdateResult(result pgconn.CommandTag) error {
	switch rowsAffected := result.RowsAffected(); rowsAffected {
	case 0:
		return blocker.ErrUserBlockNotFound
	default:
		return nil
	}
}
