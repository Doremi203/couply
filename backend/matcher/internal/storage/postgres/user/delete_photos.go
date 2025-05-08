package user

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
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
		return errors.Wrap(err, "DeletePhotos")
	}

	return nil
}
