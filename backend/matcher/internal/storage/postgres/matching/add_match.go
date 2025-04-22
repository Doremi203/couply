package matching

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (s *PgStorageMatching) AddMatch(ctx context.Context, match *matching.Match) error {
	matchSQL := `
		INSERT INTO Matches (main_user_id, chosen_user_id, approved)
		VALUES ($1, $2, $3)
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		matchSQL,
		match.MainUserID,
		match.ChosenUserID,
		match.Approved,
	)
	if err != nil {
		return fmt.Errorf("AddMatch: %w", err)
	}

	return nil
}
