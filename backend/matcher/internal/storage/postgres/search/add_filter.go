package search

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (s *PgStorageSearch) AddFilter(ctx context.Context, filter *search.Filter) error {
	filterSQL := `
		INSERT INTO filters (user_id, gender_priority, min_age, max_age, min_height, max_height, distance, goal, zodiac,
		                     education, children, alcohol, smoking, only_verified, only_premium, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		filterSQL,
		filter.UserID,
		filter.GenderPriority,
		filter.MinAge,
		filter.MaxAge,
		filter.MinHeight,
		filter.MaxHeight,
		filter.Distance,
		filter.Goal,
		filter.Zodiac,
		filter.Education,
		filter.Children,
		filter.Alcohol,
		filter.Smoking,
		filter.OnlyVerified,
		filter.OnlyPremium,
		filter.CreatedAt,
		filter.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("AddFilter: %w", err)
	}

	return nil
}
