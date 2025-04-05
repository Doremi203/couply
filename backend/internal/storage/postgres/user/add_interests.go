package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/user/interest"
	"strconv"
)

func (s *PgStorageUser) AddInterests(ctx context.Context, userID int64, interests *interest.Interest) error {
	interestSQL := `
		INSERT INTO Interests (user_id, type, value)
		VALUES ($1, $2, $3)
	`

	interestMap := map[string]string{
		"social":           toString(interests.Social),
		"sport":            toString(interests.Sport),
		"self_development": toString(interests.SelfDevelopment),
		"art":              toString(interests.Art),
		"hobby":            toString(interests.Hobby),
		"gastronomy":       toString(interests.Gastronomy),
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
			return fmt.Errorf("AddInterests: %w", err)
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
