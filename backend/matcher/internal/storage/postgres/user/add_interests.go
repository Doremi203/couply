package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

func (s *PgStorageUser) AddInterests(ctx context.Context, userID int64, interests *interest.Interest) error {
	interestSQL := `
        INSERT INTO interests (user_id, type, value)
        VALUES ($1, $2, $3)
    `

	interestGroups := map[string][]int{
		"social":           convertSlice(interests.Social),
		"sport":            convertSlice(interests.Sport),
		"self_development": convertSlice(interests.SelfDevelopment),
		"art":              convertSlice(interests.Art),
		"hobby":            convertSlice(interests.Hobby),
		"gastronomy":       convertSlice(interests.Gastronomy),
	}

	for interestType, values := range interestGroups {
		if len(values) == 0 {
			continue
		}

		for _, value := range values {
			_, err := s.txManager.GetQueryEngine(ctx).Exec(
				ctx,
				interestSQL,
				userID,
				interestType,
				value,
			)
			if err != nil {
				return fmt.Errorf("AddInterests: %w", err)
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
