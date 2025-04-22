package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) GetUser(ctx context.Context, userID int64) (*user.User, error) {
	userSQL := `
		SELECT 
			id, name, age, gender, location, bio, goal, zodiac, height, 
			education, children, alcohol, smoking, hidden, verified, 
			created_at, updated_at
		FROM Users 
		WHERE id = $1
	`

	u := &user.User{}

	err := s.txManager.GetQueryEngine(ctx).QueryRow(
		ctx,
		userSQL,
		userID,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Age,
		&u.Gender,
		&u.Location,
		&u.BIO,
		&u.Goal,
		&u.Zodiac,
		&u.Height,
		&u.Education,
		&u.Children,
		&u.Alcohol,
		&u.Smoking,
		&u.Hidden,
		&u.Verified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("GetUser: %w", err)
	}

	return u, nil
}
