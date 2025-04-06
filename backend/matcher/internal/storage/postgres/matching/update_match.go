package matching

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/matching"
)

func (s *PgStorageMatching) UpdateMatch(ctx context.Context, match *matching.Match) error {
	matchSQL := `
        UPDATE Matches 
        SET approved = $3
        WHERE main_user_id = $1 AND chosen_user_id = $2
    `

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		matchSQL,
		match.MainUserID,
		match.ChosenUserID,
		match.Approved,
	)
	if err != nil {
		return fmt.Errorf("UpdateMatch: %w", err)
	}

	return nil
}
