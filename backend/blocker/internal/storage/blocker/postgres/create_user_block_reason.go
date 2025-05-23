package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *PgStorageBlocker) CreateUserBlockReason(ctx context.Context, blockID uuid.UUID, reason blocker.ReportReason) error {
	query, args, err := buildCreateUserBlockReasonQuery(blockID, reason)
	if err != nil {
		return errors.Wrapf(err, "buildCreateUserBlockReasonQuery with %v", errors.Token("block_id", blockID))
	}

	if err = executeCreateUserBlockReasonQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return errors.Wrapf(err, "executeCreateUserBlockReasonQuery with %v", errors.Token("block_id", blockID))
	}

	return nil
}

func buildCreateUserBlockReasonQuery(blockID uuid.UUID, reason blocker.ReportReason) (string, []any, error) {
	query, args, err := sq.Insert(userBlockReasonsTableName).
		Columns(userBlockReasonsColumns...).
		Values(blockID, reason).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateUserBlockReasonQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}
