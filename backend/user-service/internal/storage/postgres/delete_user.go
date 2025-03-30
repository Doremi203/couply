package postgres

import (
	"context"
	"fmt"
)

func (s *PgStorage) DeleteUser(ctx context.Context, id int64) error {
	userSQL := `
		DELETE FROM Users
		WHERE id = $1
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, userSQL, id)
	if err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}

	return nil
}
