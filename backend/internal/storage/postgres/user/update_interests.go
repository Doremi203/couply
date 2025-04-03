package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/Couply/backend/internal/domain/user/interest"
)

func (s *PgStorageUser) UpdateInterests(ctx context.Context, userID int64, interests *interest.Interest) error {
	deleteSQL := `DELETE FROM Interests WHERE user_id = $1`
	_, err := s.txManager.GetQueryEngine(ctx).Exec(ctx, deleteSQL, userID)
	if err != nil {
		return fmt.Errorf("UpdateInterests: %w", err)
	}

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
			_, err := s.txManager.GetQueryEngine(ctx).Exec(
				ctx,
				interestSQL,
				userID,
				interestType,
				value,
			)
			if err != nil {
				return fmt.Errorf("UpdateInterests: %w", err)
			}
		}
	}

	return nil
}
