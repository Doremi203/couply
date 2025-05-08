package search

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
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
		filter.GetGenderPriority(),
		filter.GetMinAge(),
		filter.GetMaxAge(),
		filter.GetMinHeight(),
		filter.GetMaxHeight(),
		filter.GetDistance(),
		filter.GetGoal(),
		filter.GetZodiac(),
		filter.GetEducation(),
		filter.GetChildren(),
		filter.GetAlcohol(),
		filter.GetSmoking(),
		filter.GetOnlyVerified(),
		filter.GetOnlyPremium(),
		filter.GetUpdatedAt(),
		filter.GetUserID(),
	)
	if err != nil {
		return errors.Wrap(err, "UpdateFilter")
	}

	return nil
}
