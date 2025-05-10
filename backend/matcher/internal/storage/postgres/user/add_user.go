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
			user.GetID(),
			user.GetName(),
			user.GetAge(),
			user.GetGender(),
			user.GetLatitude(),
			user.GetLongitude(),
			user.GetBIO(),
			user.GetGoal(),
			user.GetZodiac(),
			user.GetHeight(),
			user.GetEducation(),
			user.GetChildren(),
			user.GetAlcohol(),
			user.GetSmoking(),
			user.GetIsHidden(),
			user.GetIsVerified(),
			user.GetIsPremium(),
			user.GetIsBlocked(),
			user.GetCreatedAt(),
			user.GetUpdatedAt(),
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
