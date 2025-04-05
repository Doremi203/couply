package matching

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/matching"
	"github.com/georgysavva/scany/pgxscan"
)

func (s *PgStorageMatching) FetchMatches(ctx context.Context, userID int64, limit, offset int32) ([]*matching.Match, error) {
	matchSQL := `
		SELECT *
		FROM Matches
		WHERE (main_user_id = $1 OR chosen_user_id = $1) AND approved = $2
		LIMIT $3 OFFSET $4
	`

	tx := s.txManager.GetQueryEngine(ctx)

	var matches []*matching.Match

	err := pgxscan.Select(
		ctx,
		tx,
		&matches,
		matchSQL,
		userID,
		true,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("FetchMatches: %w", err)
	}

	return matches, nil
}
