package postgres

import (
	"context"
	"fmt"

	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (s *PgStorage) AddUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userSQL := `
		INSERT INTO Users (name, age, gender, location, bio, goal, zodiac, height, education, children, alcohol,
		                   smoking, hidden, verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		RETURNING id
	`

	var userID int64
	err := s.txManager.GetQueryEngine(ctx).QueryRow(ctx, userSQL,
		user.Name, user.Age, user.Gender, user.Location, user.BIO, user.Goal, user.Zodiac, user.Height, user.Education,
		user.Children, user.Alcohol, user.Smoking, user.Hidden, user.Verified, user.CreatedAt, user.UpdatedAt,
	).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("AddUser: %w", err)
	}

	user.ID = userID

	return user, nil
}
