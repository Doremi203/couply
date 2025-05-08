package user

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

func (s *PgStorageUser) DeleteUser(ctx context.Context, id int64) error {
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
		return errors.Wrap(err, "DeleteUser")
	}

	return nil
}
