package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

type GetFilterOptions struct {
	UserID uuid.UUID
}

func (s *PgStorageSearch) GetFilter(ctx context.Context, opts GetFilterOptions) (*search.Filter, error) {
	query, args, err := buildGetFilterQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetFilterQuery with %v", errors.Token("options", opts))
	}

	filter, err := executeGetFilterQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetFilterQuery with %v", errors.Token("options", opts))
	}

	return filter, nil
}

func buildGetFilterQuery(opts GetFilterOptions) (string, []any, error) {
	query, args, err := sq.Select(filtersColumns...).
		From(filtersTableName).
		Where(sq.Eq{userIdColumnName: opts.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetFilterQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (*search.Filter, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	filter, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[search.Filter])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(search.ErrFilterNotFound, "pgx.CollectExactlyOneRow")
		}
		return nil, errors.Wrap(err, "pgx.CollectExactlyOneRow")
	}

	return filter, nil
}
