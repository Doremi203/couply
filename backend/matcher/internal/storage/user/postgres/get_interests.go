package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/common/interest"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type dbInterest struct {
	userID       uuid.UUID `db:"user_id"`
	interestType string    `db:"type"`
	value        int       `db:"value"`
}

type GetInterestsOptions struct {
	UserID uuid.UUID
}

func (s *PgStorageUser) GetInterests(ctx context.Context, opts GetInterestsOptions) (*interest.Interest, error) {
	query, args, err := buildGetInterestsQuery(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "buildGetInterestsQuery with %v", errors.Token("options", opts))
	}

	interestsFromDB, err := executeGetInterestsQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return nil, errors.Wrapf(err, "executeGetInterestsQuery with %v", errors.Token("options", opts))
	}

	userInterests := interest.NewInterest()
	for _, i := range interestsFromDB {
		if err = mapInterestValue(userInterests, i); err != nil {
			return nil, errors.Wrapf(err, "mapInterestValue with %v", errors.Token("options", opts))
		}
	}

	return userInterests, nil
}

func buildGetInterestsQuery(opts GetInterestsOptions) (string, []any, error) {
	query, args, err := sq.Select(interestsColumns...).
		From(interestsTableName).
		Where(sq.Eq{userIDColumnName: opts.UserID}).
		OrderBy(typeColumnName).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeGetInterestsQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) ([]dbInterest, error) {
	rows, err := queryEngine.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	interests, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbInterest])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Wrap(interest.ErrInterestsNotFound, "query")
		}
		return nil, errors.Wrap(err, "pgx.CollectRows")
	}

	return interests, nil
}

func mapInterestValue(i *interest.Interest, dbInterest dbInterest) error {
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
