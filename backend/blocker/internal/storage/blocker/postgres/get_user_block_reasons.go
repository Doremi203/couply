package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/storage"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type GetUserBlockReasonsOptions struct {
	BlockID uuid.UUID
}

func (s *PgStorageBlocker) GetUserBlockReasons(ctx context.Context, opts GetUserBlockReasonsOptions) ([]blocker.ReportReason, error) {
	query, args, err := buildGetUserBlockReasonsQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetUserBlockReasonsQuery with %v", errors.Token("options", opts))
	}

	reasons, err := executeGetUserBlockReasonsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetUserBlockReasonsQuery with %v", errors.Token("options", opts))
	}

	return reasons, nil
}

func buildGetUserBlockReasonsQuery(opts GetUserBlockReasonsOptions) (string, []any, error) {
	query, args, err := sq.Select("reason").
		From("user_block_reasons").
		Where(sq.Eq{"block_id": opts.BlockID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetUserBlockReasonsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]blocker.ReportReason, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	reasons, err := pgx.CollectRows(rows, pgx.RowTo[blocker.ReportReason])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(blocker.ErrUserBlockReasonsNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return reasons, nil
}
