package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetMatchOptions struct {
	FirstUserID  uuid.UUID
	SecondUserID uuid.UUID
}

func (s *PgStorageMatching) GetMatch(ctx context.Context, opts GetMatchOptions) (*matching.Match, error) {
	const query = `
		SELECT * FROM matches
		WHERE (first_user_id = $1 AND second_user_id = $2) OR (first_user_id = $2 AND second_user_id = $1)
	`

	like, err := executeGetMatchQuery(ctx, s.txManager.GetQueryEngine(ctx), query, []any{opts.FirstUserID, opts.SecondUserID})
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetMatchQuery with %v", errors.Token("options", opts))
	}

	return like, nil
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
