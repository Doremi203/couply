package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/user"
)

func (s *PgStorageUser) GetUser(ctx context.Context, userID int64) (*user.User, error) {
	userSQL := `
		SELECT *
		FROM Users
		WHERE id = $1
	`

	tx := s.txManager.GetQueryEngine(ctx)

	var user *user.User

	err := pgxscan.Get(
		ctx,
		tx,
		&user,
		userSQL,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("GetUser: %w", err)
	}

	return user, nil
}
