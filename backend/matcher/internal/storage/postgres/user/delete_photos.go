package user

import (
	"context"
	"fmt"
)

func (s *PgStorageUser) DeletePhotos(ctx context.Context, id int64) error {
	photoSQL := `
		DELETE FROM photos
		WHERE user_id = $1
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		photoSQL,
		id,
	)
	if err != nil {
		return fmt.Errorf("DeletePhotos: %w", err)
	}

	return nil
}
