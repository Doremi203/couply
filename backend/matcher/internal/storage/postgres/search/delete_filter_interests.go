package search

import (
	"context"
	"fmt"
)

func (s *PgStorageSearch) DeleteFilterInterests(ctx context.Context, id int64) error {
	interestsSQL := `
		DELETE FROM filter_interests
		WHERE user_id = $1
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		interestsSQL,
		id,
	)
	if err != nil {
		return fmt.Errorf("DeleteFilterInterests: %w", err)
	}

	return nil
}
