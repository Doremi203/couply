package search

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

func (s *PgStorageSearch) AddFilterInterests(ctx context.Context, userID int64, filterInterests *interest.Interest) error {
	filterInterestsSQL := `
        INSERT INTO filter_interests (user_id, type, value)
        VALUES ($1, $2, $3)
    `

	interestGroups := map[string][]int{
		"social":           convertSlice(filterInterests.Social),
		"sport":            convertSlice(filterInterests.Sport),
		"self_development": convertSlice(filterInterests.SelfDevelopment),
		"art":              convertSlice(filterInterests.Art),
		"hobby":            convertSlice(filterInterests.Hobby),
		"gastronomy":       convertSlice(filterInterests.Gastronomy),
	}

	for interestType, values := range interestGroups {
		if len(values) == 0 {
			continue
		}

		for _, value := range values {
			_, err := s.txManager.GetQueryEngine(ctx).Exec(
				ctx,
				filterInterestsSQL,
				userID,
				interestType,
				value,
			)
			if err != nil {
				return fmt.Errorf("AddFilterInterests: %w", err)
			}
		}
	}

	return nil
}

func convertSlice[T ~int](slice []T) []int {
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		result = append(result, int(v))
	}
	return result
}
