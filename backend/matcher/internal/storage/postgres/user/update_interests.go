package user

import (
	"context"
	"fmt"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

func (s *PgStorageUser) UpdateInterests(ctx context.Context, userID int64, interests *interest.Interest) error {
	interestSQL := `
        UPDATE interests
        SET value = $3
        WHERE user_id = $1 AND type = $2
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
			return fmt.Errorf("UpdateInterests: %w", err)
		}
	}

	return nil
}
