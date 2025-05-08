package matching

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/matching"
)

func (s *PgStorageMatching) AddMatch(ctx context.Context, match *matching.Match) error {
	matchSQL := `
		INSERT INTO matches (main_user_id, chosen_user_id, approved)
		VALUES ($1, $2, $3)
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		matchSQL,
		match.GetMainUserID(),
		match.GetChosenUserID(),
		match.GetApproved(),
	)
	if err != nil {
		return errors.Wrap(err, "AddMatch")
	}

	return nil
}
