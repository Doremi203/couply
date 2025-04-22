package matching

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
	"github.com/jackc/pgx/v5"
)

func (s *PgStorageMatching) FetchIncomingMatches(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error) {
	matchSQL := `
		SELECT *
		FROM Matches
		WHERE chosen_user_id = $1 AND approved = $2
		LIMIT $3 OFFSET $4
	`

	tx := s.txManager.GetQueryEngine(ctx)

	rows, err := tx.Query(ctx, matchSQL, userID, false, limit, offset)
	if err != nil {
		return nil, errors.WrapFail(err, "fetch incoming matches")
	}
	defer rows.Close()

	matches, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[matching.Match])
	if err != nil {
		return nil, errors.WrapFail(err, "unmarshal incoming matches to struct")
	}

	return matches, nil
}
