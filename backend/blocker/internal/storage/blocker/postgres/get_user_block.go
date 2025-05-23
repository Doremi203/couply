package postgres

import (
	"context"
	"github.com/Doremi203/couply/backend/blocker/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetUserBlockOptions struct {
	BlockID       uuid.UUID
	UserID        uuid.UUID
	AcceptedBlock bool
	ForUpdate     bool
}

func (s *PgStorageBlocker) GetUserBlock(ctx context.Context, opts GetUserBlockOptions) (*blocker.UserBlock, error) {
	query, args, err := buildGetUserBlockQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetUserBlockQuery with %v", errors.Token("options", opts))
	}

	block, err := executeGetUserBlockQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetUserBlockQuery with %v", errors.Token("options", opts))
	}

	return block, nil
}

func buildGetUserBlockQuery(options GetUserBlockOptions) (string, []any, error) {
	sb := sq.Select(userBlocksColumns...).
		From(userBlocksTableName)

	if options.BlockID != uuid.Nil {
		sb = sb.Where(sq.Eq{idColumn: options.BlockID})
	}
	if options.AcceptedBlock {
		sb = sb.Where(sq.Eq{blockedIdColumn: options.UserID}).
			Where(sq.Eq{statusColumn: blocker.BlockStatusAccepted})
	}
	if options.ForUpdate {
		sb = sb.Suffix("FOR UPDATE")
	}

	return sb.PlaceholderFormat(sq.Dollar).ToSql()
}

func executeGetUserBlockQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (*blocker.UserBlock, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	block, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[blocker.UserBlock])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(blocker.ErrUserBlockNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectExactlyOneRow")
	}

	return block, nil
}
