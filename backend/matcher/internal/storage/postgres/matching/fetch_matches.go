package matching

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/jackc/pgx/v5"
)

func (s *PgStorageMatching) FetchMatches(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error) {
	matchSQL := `
		SELECT *
		FROM matches
		WHERE (main_user_id = $1 OR chosen_user_id = $1) AND approved = $2
		LIMIT $3 OFFSET $4
	`

	tx := s.txManager.GetQueryEngine(ctx)

	rows, err := tx.Query(ctx, matchSQL, userID, true, limit, offset)
	if err != nil {
		return nil, errors.WrapFail(err, "fetch matches")
	}
	defer rows.Close()

	matches, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[matching.Match])
	if err != nil {
		return nil, errors.WrapFail(err, "unmarshal matches to struct")
	}

	return matches, nil
}
