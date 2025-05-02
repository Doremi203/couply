package search

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	"strconv"
)

func (s *PgStorageSearch) AddFilterInterests(ctx context.Context, userID int64, filterInterests *interest.Interest) error {
	interestSQL := `
		INSERT INTO interests (user_id, type, value)
		VALUES ($1, $2, $3)
	`

	interestMap := map[string]string{
		"social":           toString(filterInterests.Social),
		"sport":            toString(filterInterests.Sport),
		"self_development": toString(filterInterests.SelfDevelopment),
		"art":              toString(filterInterests.Art),
		"hobby":            toString(filterInterests.Hobby),
		"gastronomy":       toString(filterInterests.Gastronomy),
	}

	for interestType, value := range interestMap {
		_, err := s.txManager.GetQueryEngine(ctx).Exec(
			ctx,
			interestSQL,
			userID,
			interestType,
			value,
		)
		if err != nil {
			return fmt.Errorf("AddFilterInterests: %w", err)
		}
	}

	return nil
}

// вспомогательная функция для преобразования срезов enum в срезы int
func toString[T ~int](slice []T) string {
	result := ""
	for _, v := range slice {
		result += strconv.Itoa(int(v))
	}
	return result
}
