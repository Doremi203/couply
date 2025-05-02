package search

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (s *PgStorageSearch) UpdateFilter(ctx context.Context, filter *search.Filter) error {
	filterSQL := `
		UPDATE filters 
		SET gender_priority = $1, min_age = $2, max_age = $3, min_height = $4, max_height = $5, distance = $6,
		    goal = $7, zodiac = $8, education = $9, children = $10, alcohol = $11, smoking = $12, only_verified = $13,
		    only_premium = $14, updated_at = $15
		WHERE user_id = $16
	`

	_, err := s.txManager.GetQueryEngine(ctx).Exec(
		ctx,
		filterSQL,
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
		filter.UpdatedAt,
		filter.UserID,
	)
	if err != nil {
		return fmt.Errorf("UpdateFilter: %w", err)
	}

	return nil
}
