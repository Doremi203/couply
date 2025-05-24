package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageUser) UpdateUser(ctx context.Context, user *user.User) (*user.User, error) {
	query, args, err := sq.Update("users").
		Set("name", user.Name).
		Set("age", user.Age).
		Set("gender", user.Gender).
		Set("latitude", user.Latitude).
		Set("longitude", user.Longitude).
		Set("bio", user.BIO).
		Set("goal", user.Goal).
		Set("zodiac", user.Zodiac).
		Set("height", user.Height).
		Set("education", user.Education).
		Set("children", user.Children).
		Set("alcohol", user.Alcohol).
		Set("smoking", user.Smoking).
		Set("is_hidden", user.IsHidden).
		Set("is_verified", user.IsVerified).
		Set("is_premium", user.IsPremium).
		Set("is_blocked", user.IsBlocked).
		Set("updated_at", user.UpdatedAt).
		Where(sq.Eq{"id": user.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	result, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return nil, ErrUserNotFound
	}

	return user, nil
}
