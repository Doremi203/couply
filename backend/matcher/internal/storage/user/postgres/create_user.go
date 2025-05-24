package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/Doremi203/couply/backend/matcher/internal/storage"

	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	sq "github.com/Masterminds/squirrel"
)

func (s *PgStorageUser) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	query, args, err := buildCreateUserQuery(user)
	if err != nil {
		return nil, errors.Wrapf(err, "buildCreateUserQuery with %v", errors.Token("user", user))
	}

	if err = executeCreateUserQuery(ctx, s.txManager.GetQueryEngine(ctx), query, args); err != nil {
		return nil, errors.Wrapf(err, "executeCreateUserQuery with %v", errors.Token("user", user))
	}

	return user, nil
}

func buildCreateUserQuery(user *user.User) (string, []any, error) {
	query, args, err := sq.Insert(usersTableName).
		Columns(usersColumns...).
		Values(user.ID, user.Name, user.Age, user.Gender, user.Latitude, user.Longitude, user.BIO, user.Goal,
			user.Zodiac, user.Height, user.Education, user.Children, user.Alcohol, user.Smoking, user.IsHidden,
			user.IsVerified, user.IsPremium, user.IsBlocked, user.CreatedAt, user.UpdatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	return query, args, err
}

func executeCreateUserQuery(ctx context.Context, queryEngine storage.QueryEngine, query string, args []any) error {
	_, err := queryEngine.Exec(ctx, query, args...)
	if err != nil {
		if postgres.IsUniqueViolationError(err) {
			return user.ErrDuplicateUser
		}
		return errors.Wrap(err, "exec")
	}
	return nil
}
