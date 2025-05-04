package search

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
)

func (s *PgStorageSearch) GetFilter(ctx context.Context, userID int64) (*search.Filter, error) {
	filterSQL := `
		SELECT * 
		FROM filters
		WHERE user_id = $1
	`

	filter := &search.Filter{}

	err := s.txManager.GetQueryEngine(ctx).QueryRow(
		ctx,
		filterSQL,
		userID,
	).Scan(
		&filter.UserID,
		&filter.GenderPriority,
		&filter.MinAge,
		&filter.MaxAge,
		&filter.MinHeight,
		&filter.MaxHeight,
		&filter.Distance,
		&filter.Goal,
		&filter.Zodiac,
		&filter.Education,
		&filter.Children,
		&filter.Alcohol,
		&filter.Smoking,
		&filter.OnlyVerified,
		&filter.OnlyPremium,
		&filter.CreatedAt,
		&filter.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("GetFilter: %w", err)
	}
	return filter, nil
}
