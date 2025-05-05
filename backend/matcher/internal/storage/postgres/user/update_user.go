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
		user.GetName(),
		user.GetAge(),
		user.GetGender(),
		user.GetLocation(),
		user.GetBIO(),
		user.GetGoal(),
		user.GetZodiac(),
		user.GetHeight(),
		user.GetEducation(),
		user.GetChildren(),
		user.GetAlcohol(),
		user.GetSmoking(),
		user.GetHidden(),
		user.GetVerified(),
		user.GetUpdatedAt(),
		user.GetID(),
	)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser: %w", err)
	}

	return user, nil
}
