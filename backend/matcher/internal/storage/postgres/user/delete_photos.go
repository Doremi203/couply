package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *PgStorageUser) DeletePhotos(ctx context.Context, id uuid.UUID) error {
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
