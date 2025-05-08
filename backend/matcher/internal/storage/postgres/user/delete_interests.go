package user

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

func (s *PgStorageUser) DeleteInterests(ctx context.Context, id int64) error {
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
		return errors.Wrap(err, "DeleteInterests")
	}

	return nil
}
