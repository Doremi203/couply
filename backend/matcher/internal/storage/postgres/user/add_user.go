package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageUser) AddUser(ctx context.Context, user *user.User) (*user.User, error) {
	query, args, err := sq.Insert("users").
		Columns(
			"id", "name", "age", "gender", "latitude", "longitude", "bio", "goal", "zodiac",
			"height", "education", "children", "alcohol", "smoking", "is_hidden",
			"is_verified", "is_premium", "is_blocked", "created_at", "updated_at",
		).
		Values(
			user.ID,
			user.Name,
			user.Age,
			user.Gender,
			user.Latitude,
			user.Longitude,
			user.BIO,
			user.Goal,
			user.Zodiac,
			user.Height,
			user.Education,
			user.Children,
			user.Alcohol,
			user.Smoking,
			user.IsHidden,
			user.IsVerified,
			user.IsPremium,
			user.IsBlocked,
			user.CreatedAt,
			user.UpdatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return user, nil
}
