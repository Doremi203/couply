package user

import (
	"context"
	"fmt"
)

func (s *PgStorageUser) DeleteInterests(ctx context.Context, id int64) error {
	interestsSQL := `
		DELETE FROM Interests
		WHERE user_id = $1
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		interestsSQL,
		id,
	)
	if err != nil {
		return fmt.Errorf("DeleteInterests: %w", err)
	}

	return nil
}
