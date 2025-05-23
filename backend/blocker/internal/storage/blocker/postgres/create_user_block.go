package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	pgerrors "github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/blocker/internal/storage"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageBlocker) CreateUserBlock(ctx context.Context, block *blocker.UserBlock) error {
	query, args, err := buildCreateUserBlockQuery(block)
	if err != nil {
		return errors.Wrapf(err, "buildCreateUserBlockQuery with %v", errors.Token("block", block))
	}

	if err = executeCreateUserBlockQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreateUserBlockQuery with %v", errors.Token("block", block))
	}

	return nil
}

func buildCreateUserBlockQuery(block *blocker.UserBlock) (string, []any, error) {
	query, args, err := sq.Insert(userBlocksTableName).
		Columns(userBlocksColumns...).
		Values(block.ID, block.BlockedID, block.Message, block.CreatedAt, block.Status).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateUserBlockQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if pgerrors.IsUniqueViolationError(err) {
			return errors.Wrap(
				blocker.ErrDuplicateUserBlock,
				"exec",
			)
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
