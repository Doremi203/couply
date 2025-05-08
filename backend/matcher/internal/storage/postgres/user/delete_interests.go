package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *PgStorageUser) DeleteInterests(ctx context.Context, id uuid.UUID) error {
	interestsSQL := `
		DELETE FROM interests
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
