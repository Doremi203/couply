package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (s *PgStorageUser) UpdateUser(ctx context.Context, user *user.User) (*user.User, error) {
	userSQL := `
        UPDATE users 
        SET name = $1, age = $2, gender = $3, location = $4, bio = $5, goal = $6, zodiac = $7, 
            height = $8, education = $9, children = $10, alcohol = $11, smoking = $12, 
            hidden = $13, verified = $14, updated_at = $15
        WHERE id = $16
    `

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		userSQL,
		user.Name,
		user.Age,
		user.Gender,
		user.Location,
		user.BIO,
		user.Goal,
		user.Zodiac,
		user.Height,
		user.Education,
		user.Children,
		user.Alcohol,
		user.Smoking,
		user.Hidden,
		user.Verified,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser: %w", err)
	}

	return user, nil
}
