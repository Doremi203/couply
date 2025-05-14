package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageUser) UpdateUser(ctx context.Context, user *user.User) (*user.User, error) {
	query, args, err := sq.Update("users").
		Set("name", user.GetName()).
		Set("age", user.GetAge()).
		Set("gender", user.GetGender()).
		Set("latitude", user.GetLatitude()).
		Set("longitude", user.GetLongitude()).
		Set("bio", user.GetBIO()).
		Set("goal", user.GetGoal()).
		Set("zodiac", user.GetZodiac()).
		Set("height", user.GetHeight()).
		Set("education", user.GetEducation()).
		Set("children", user.GetChildren()).
		Set("alcohol", user.GetAlcohol()).
		Set("smoking", user.GetSmoking()).
		Set("is_hidden", user.GetIsHidden()).
		Set("is_verified", user.GetIsVerified()).
		Set("is_premium", user.GetIsPremium()).
		Set("is_blocked", user.GetIsBlocked()).
		Set("updated_at", user.GetUpdatedAt()).
		Where(sq.Eq{"id": user.GetID()}).
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
