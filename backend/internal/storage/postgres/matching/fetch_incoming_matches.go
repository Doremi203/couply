package matching

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/matching"
	"github.com/Doremi203/Couply/backend/internal/storage/postgres/user"
)

func (s *PgStorageMatching) FetchIncomingMatches(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error) {
	matchSQL := `
		SELECT *
		FROM Matches
		WHERE chosen_user_id = $1 AND approved = $2
		LIMIT $3 OFFSET $4
	`

	tx := s.txManager.GetQueryEngine(ctx)

	var matches []*matching.Match

	err := pgxscan.Get(
		ctx,
		tx,
		&matches,
		matchSQL,
		userID,
		false,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("FetchIncomingMatches: %w", err)
	}

	return matches, nil
}
