package search

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageSearch) UpdateFilter(ctx context.Context, filter *search.Filter) error {
	query, args, err := sq.Update("filters").
		Set("gender_priority", filter.GetGenderPriority()).
		Set("min_age", filter.GetMinAge()).
		Set("max_age", filter.GetMaxAge()).
		Set("min_height", filter.GetMinHeight()).
		Set("max_height", filter.GetMaxHeight()).
		Set("min_distance_km", filter.GetMinDistanceKM()).
		Set("max_distance_km", filter.GetMaxDistanceKM()).
		Set("goal", filter.GetGoal()).
		Set("zodiac", filter.GetZodiac()).
		Set("education", filter.GetEducation()).
		Set("children", filter.GetChildren()).
		Set("alcohol", filter.GetAlcohol()).
		Set("smoking", filter.GetSmoking()).
		Set("only_verified", filter.GetOnlyVerified()).
		Set("only_premium", filter.GetOnlyPremium()).
		Set("updated_at", filter.GetUpdatedAt()).
		Where(sq.Eq{"user_id": filter.GetUserID()}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	result, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrFilterNotFound
	}

	return nil
}
