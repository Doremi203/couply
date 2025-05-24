package search

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

func (s *PgStorageSearch) AddFilterInterests(ctx context.Context, userID uuid.UUID, filterInterests *interest.Interest) error {
	interestGroups := map[string][]int{
		interest.SportDBName:             convertSlice(filterInterests.Sport),
		interest.SelfDevelopmentDBName:   convertSlice(filterInterests.SelfDevelopment),
		interest.HobbyDBName:             convertSlice(filterInterests.Hobby),
		interest.MusicDBName:             convertSlice(filterInterests.Music),
		interest.MoviesTVDBName:          convertSlice(filterInterests.MoviesTV),
		interest.FoodDrinkDBName:         convertSlice(filterInterests.FoodDrink),
		interest.PersonalityTraitsDBName: convertSlice(filterInterests.PersonalityTraits),
		interest.PetsDBName:              convertSlice(filterInterests.Pets),
	}

	for interestType, values := range interestGroups {
		if len(values) == 0 {
			continue
		}

		for _, value := range values {
			query, args, err := sq.Insert("filter_interests").
				Columns("user_id", "type", "value").
				Values(userID, interestType, value).
				PlaceholderFormat(sq.Dollar).
				ToSql()
			if err != nil {
				return fmt.Errorf("failed to build query: %w", err)
			}

			_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
			if err != nil {
				return fmt.Errorf("failed to insert filter interest: %w", err)
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
