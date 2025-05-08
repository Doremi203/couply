package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *PgStorageUser) DeleteUser(ctx context.Context, id uuid.UUID) error {
	userSQL := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		userSQL,
		id,
	)
	if err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}

	return nil
}
