package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	sq "github.com/Masterminds/squirrel"
)

type FetchMatchesOptions struct {
	UserID uuid.UUID
	Limit  uint64
	Offset uint64
}

func (s *PgStorageMatching) FetchMatches(ctx context.Context, opts FetchMatchesOptions) ([]*matching.Match, error) {
	query, args, err := buildFetchMatchesQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildFetchMatchesQuery with %v", errors.Token("options", opts))
	}

	matches, err := executeFetchMatchesQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeFetchMatchesQuery with %v", errors.Token("options", opts))
	}

	return matches, nil
}

func buildFetchMatchesQuery(opts FetchMatchesOptions) (string, []any, error) {
	query, args, err := sq.Select(matchesColumns...).
		From(matchesTableName).
		Where(sq.Or{
			sq.Eq{firstUserIDColumnName: opts.UserID},
			sq.Eq{secondUserIDColumnName: opts.UserID},
		}).
		OrderBy(createdAtColumnName + " DESC").
		Limit(opts.Limit).
		Offset(opts.Offset).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeFetchMatchesQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]*matching.Match, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	likes, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[matching.Match])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(matching.ErrMatchesNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return likes, nil
}
