package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetMatchOptions struct {
	UserID uuid.UUID
}

func (s *PgStorageMatching) GetMatch(ctx context.Context, opts GetMatchOptions) (*matching.Match, error) {
	query, args, err := buildGetMatchQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetMatchQuery with %v", errors.Token("options", opts))
	}

	like, err := executeGetMatchQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetMatchQuery with %v", errors.Token("options", opts))
	}

	return like, nil
}

func buildGetMatchQuery(opts GetMatchOptions) (string, []any, error) {
	query, args, err := sq.Select(matchesColumns...).
		From(matchesTableName).
		Where(sq.Or{
			sq.Eq{firstUserIDColumnName: opts.UserID},
			sq.Eq{secondUserIDColumnName: opts.UserID},
		}).PlaceholderFormat(sq.Dollar).
		ToSql()

	return query, args, err
}

func executeGetMatchQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (*matching.Match, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	match, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByNameLax[matching.Match])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(matching.ErrMatchNotFound, "pgx.CollectExactlyOneRow")
		}
		return nil, errors.Wrap(err, "pgx.CollectExactlyOneRow")
	}

	return match, nil
}
