package postgres

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain"
)

func (s *PgStorage) AddUser(ctx context.Context, user domain.User) error {
	const op = "AddUser"

	userSQL := `
		INSERT INTO Users (name, age, gender, location, bio, goal, zodiac, height, education, children, alcohol,
		                   smoking, hidden, verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, userSQL,
		user.Name, user.Age, user.Gender, user.Location, user.BIO, user.Goal, user.Zodiac, user.Height, user.Education,
		user.Children, user.Alcohol, user.Smoking, user.Hidden, user.Verified, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
