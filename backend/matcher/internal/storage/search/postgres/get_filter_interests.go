package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
)

type GetFilterInterestsOptions struct {
	UserID uuid.UUID
}

type dbFilterInterest struct {
	interestType string `db:"type"`
	value        int    `db:"value"`
}

func (s *PgStorageSearch) GetFilterInterests(ctx context.Context, opts GetFilterInterestsOptions) (*interest.Interest, error) {
	query, args, err := buildGetFilterInterestsQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetFilterInterestsQuery with %v", errors.Token("options", opts))
	}

	interestsFromDB, err := executeGetFilterInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetFilterInterestsQuery with %v", errors.Token("options", opts))
	}

	filterInterests := interest.NewInterest()
	for _, i := range interestsFromDB {
		if err = mapInterestValue(filterInterests, i); err != nil {
			return nil, errors.Wrapf(err, "mapInterestValue with %v", errors.Token("options", opts))
		}
	}

	return filterInterests, nil
}

func buildGetFilterInterestsQuery(opts GetFilterInterestsOptions) (string, []any, error) {
	query, args, err := sq.Select(filterInterestsColumns...).
		From(filtersTableName).
		Where(sq.Eq{userIdColumnName: opts.UserID}).
		OrderBy(typeColumnName).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetFilterInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]dbFilterInterest, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	interests, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbFilterInterest])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(interest.ErrInterestsNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return interests, nil
}

func mapInterestValue(i *interest.Interest, dbInterest dbFilterInterest) error {
	switch dbInterest.interestType {
	case interest.SportName:
		i.Sport = append(i.Sport, interest.Sport(dbInterest.value))
	case interest.SelfDevelopmentName:
		i.SelfDevelopment = append(i.SelfDevelopment, interest.SelfDevelopment(dbInterest.value))
	case interest.HobbyName:
		i.Hobby = append(i.Hobby, interest.Hobby(dbInterest.value))
	case interest.MusicName:
		i.Music = append(i.Music, interest.Music(dbInterest.value))
	case interest.MoviesTVName:
		i.MoviesTV = append(i.MoviesTV, interest.MoviesTV(dbInterest.value))
	case interest.FoodDrinkName:
		i.FoodDrink = append(i.FoodDrink, interest.FoodDrink(dbInterest.value))
	case interest.PersonalityTraitsName:
		i.PersonalityTraits = append(i.PersonalityTraits, interest.PersonalityTraits(dbInterest.value))
	case interest.PetsName:
		i.Pets = append(i.Pets, interest.Pets(dbInterest.value))
	default:
		return errors.Wrapf(interest.ErrInterestsNotFound, " %v", dbInterest.interestType)
	}
	return nil
}
