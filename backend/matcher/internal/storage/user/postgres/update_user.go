package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageUser) UpdateUser(ctx context.Context, userForUpdate *user.User) error {
	query, args, err := buildUpdateUserQuery(userForUpdate)
	if err != nil {
		return errors.Wrapf(err, "buildUpdateUserQuery with %v", errors.Token("user_id", userForUpdate.ID))
	}

	result, err := executeUpdateUserQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args)
	if err != nil {
		return errors.Wrapf(err, "executeUpdateUserQuery with %v", errors.Token("user_id", userForUpdate.ID))
	}

	if result.RowsAffected() == 0 {
		return user.ErrUserNotFound
	}

	return nil
}

func buildUpdateUserQuery(user *user.User) (string, []any, error) {
	query, args, err := sq.Update(usersTableName).
		Set(nameColumnName, user.Name).
		Set(ageColumnName, user.Age).
		Set(genderColumnName, user.Gender).
		Set(latitudeColumnName, user.Latitude).
		Set(longitudeColumnName, user.Longitude).
		Set(bioColumnName, user.BIO).
		Set(goalColumnName, user.Goal).
		Set(zodiacColumnName, user.Zodiac).
		Set(heightColumnName, user.Height).
		Set(educationColumnName, user.Education).
		Set(childrenColumnName, user.Children).
		Set(alcoholColumnName, user.Alcohol).
		Set(smokingColumnName, user.Smoking).
		Set(isHiddenColumnName, user.IsHidden).
		Set(isVerifiedColumnName, user.IsVerified).
		Set(isPremiumColumnName, user.IsPremium).
		Set(isBlockedColumnName, user.IsBlocked).
		Set(updatedAtColumnName, user.UpdatedAt).
		Where(sq.Eq{idColumnName: user.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeUpdateUserQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) (pgconn.CommandTag, error) {
	result, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, errors.Wrap(err, "exec")
	}
	return result, nil
}
