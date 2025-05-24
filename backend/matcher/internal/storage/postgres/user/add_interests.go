package user

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

func (s *PgStorageUser) AddInterests(ctx context.Context, userID uuid.UUID, interests *interest.Interest) error {
	interestGroups := map[string][]int{
		interest.SportDBName:             convertSlice(interests.Sport),
		interest.SelfDevelopmentDBName:   convertSlice(interests.SelfDevelopment),
		interest.HobbyDBName:             convertSlice(interests.Hobby),
		interest.MusicDBName:             convertSlice(interests.Music),
		interest.MoviesTVDBName:          convertSlice(interests.MoviesTV),
		interest.FoodDrinkDBName:         convertSlice(interests.FoodDrink),
		interest.PersonalityTraitsDBName: convertSlice(interests.PersonalityTraits),
		interest.PetsDBName:              convertSlice(interests.Pets),
	}

	for interestType, values := range interestGroups {
		if len(values) == 0 {
			continue
		}

		for _, value := range values {
			query, args, err := sq.Insert("interests").
				Columns("user_id", "type", "value").
				Values(userID, interestType, value).
				PlaceholderFormat(sq.Dollar).
				ToSql()
			if err != nil {
				return fmt.Errorf("failed to build query: %w", err)
			}

			_, err = s.txManager.GetQueryEngine(ctx).Exec(ctx, query, args...)
			if err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) && pgErr.Code == "23505" {
					return ErrDuplicateInterest
				}
				return fmt.Errorf("failed to insert interest: %w", err)
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
