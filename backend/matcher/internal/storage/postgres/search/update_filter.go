package search

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageSearch) UpdateFilter(ctx context.Context, filter *search.Filter) error {
	query, args, err := sq.Update("filters").
		Set("gender_priority", filter.GenderPriority).
		Set("min_age", filter.MinAge).
		Set("max_age", filter.MaxAge).
		Set("min_height", filter.MinHeight).
		Set("max_height", filter.MaxHeight).
		Set("min_distance_km", filter.MinDistanceKM).
		Set("max_distance_km", filter.MaxDistanceKM).
		Set("goal", filter.Goal).
		Set("zodiac", filter.Zodiac).
		Set("education", filter.Education).
		Set("children", filter.Children).
		Set("alcohol", filter.Alcohol).
		Set("smoking", filter.Smoking).
		Set("only_verified", filter.OnlyVerified).
		Set("only_premium", filter.OnlyPremium).
		Set("updated_at", filter.UpdatedAt).
		Where(sq.Eq{"user_id": filter.UserID}).
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
