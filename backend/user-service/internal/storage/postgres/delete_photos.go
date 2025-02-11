package postgres

import (
	"context"
	"fmt"
)

func (s *PgStorage) DeletePhotos(ctx context.Context, userID int64) error {
	const op = "DeletePhotos"

	sql := `DELETE FROM Photos WHERE user_id = $1`
	_, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, sql, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
