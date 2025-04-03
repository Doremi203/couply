package matching

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/matching"
)

func (s *PgStorageMatching) FetchOutgoingMatches(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error) {
	matchSQL := `
		SELECT *
		FROM Matches
		WHERE main_user_id = $1 AND approved = $2
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
		return nil, fmt.Errorf("FetcOutgoingMatches: %w", err)
	}

	return matches, nil
}
