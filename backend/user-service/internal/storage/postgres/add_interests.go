package postgres

import (
	"context"
	"fmt"

	"github.com/Doremi203/Couply/backend/internal/domain/interest"
)

func (s *PgStorage) AddInterests(ctx context.Context, userID int64, interests *interest.Interest) error {
	interestSQL := `
		INSERT INTO Interests (user_id, type, value)
		VALUES ($1, $2, $3)
	`

	interestMap := map[string][]int{
		"social":           toIntSlice(interests.Social),
		"sport":            toIntSlice(interests.Sport),
		"self_development": toIntSlice(interests.SelfDevelopment),
		"art":              toIntSlice(interests.Art),
		"hobby":            toIntSlice(interests.Hobby),
		"gastronomy":       toIntSlice(interests.Gastronomy),
	}

	for interestType, values := range interestMap {
		for _, value := range values {
			_, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, interestSQL, userID, interestType, value)
			if err != nil {
				return fmt.Errorf("AddInterests: %w", err)
			}
		}
	}

	return nil
}

// вспомогательная функция для преобразования срезов enum в срезы int
func toIntSlice[T ~int](slice []T) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = int(v)
	}
	return result
}
