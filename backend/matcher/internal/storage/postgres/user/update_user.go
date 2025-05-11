package user

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageUser) UpdateUser(ctx context.Context, user *user.User, updateMask *fieldmaskpb.FieldMask) (*user.User, error) {
	updateBuilder := sq.Update("users").
		Set("updated_at", user.GetUpdatedAt()).
		Where(sq.Eq{"id": user.GetID()}).
		PlaceholderFormat(sq.Dollar)

	if updateMask == nil {
		updateBuilder = updateBuilder.
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
			Set("is_blocked", user.GetIsBlocked())
	} else {
		for _, path := range updateMask.GetPaths() {
			switch path {
			case "name":
				updateBuilder = updateBuilder.Set("name", user.GetName())
			case "age":
				updateBuilder = updateBuilder.Set("age", user.GetAge())
			case "gender":
				updateBuilder = updateBuilder.Set("gender", user.GetGender())
			case "latitude":
				updateBuilder = updateBuilder.Set("latitude", user.GetLatitude())
			case "longitude":
				updateBuilder = updateBuilder.Set("longitude", user.GetLongitude())
			case "bio":
				updateBuilder = updateBuilder.Set("bio", user.GetBIO())
			case "goal":
				updateBuilder = updateBuilder.Set("goal", user.GetGoal())
			case "zodiac":
				updateBuilder = updateBuilder.Set("zodiac", user.GetZodiac())
			case "height":
				updateBuilder = updateBuilder.Set("height", user.GetHeight())
			case "education":
				updateBuilder = updateBuilder.Set("education", user.GetEducation())
			case "children":
				updateBuilder = updateBuilder.Set("children", user.GetChildren())
			case "alcohol":
				updateBuilder = updateBuilder.Set("alcohol", user.GetAlcohol())
			case "smoking":
				updateBuilder = updateBuilder.Set("smoking", user.GetSmoking())
			case "is_hidden":
				updateBuilder = updateBuilder.Set("is_hidden", user.GetIsHidden())
			case "is_verified":
				updateBuilder = updateBuilder.Set("is_verified", user.GetIsVerified())
			case "is_premium":
				updateBuilder = updateBuilder.Set("is_premium", user.GetIsPremium())
			case "is_blocked":
				updateBuilder = updateBuilder.Set("is_blocked", user.GetIsBlocked())
			}
		}
	}

	query, args, err := updateBuilder.ToSql()
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
