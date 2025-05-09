package userpostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

func NewRepo(
	db postgres.Client,
) *repo {
	return &repo{
		db: db,
	}
}

type repo struct {
	db postgres.Client
}

func (r *repo) Create(ctx context.Context, u user.User) error {
	const query = `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
		ON CONFLICT (email) DO NOTHING;
	`
	res, err := r.db.Exec(ctx, query, u.ID, u.Email, u.Password)
	if err != nil {
		return errors.WrapFail(err, "save user")
	}
	if res.RowsAffected() == 0 {
		return errors.Wrapf(user.ErrAlreadyExists, "user with %v", errors.Token("email", u.Email))
	}

	return nil
}

func (r *repo) UpdatePhone(ctx context.Context, userID user.ID, phone user.Phone) error {
	const query = `
		UPDATE users SET phone = $2
		WHERE id = $1
	`
	res, err := r.db.Exec(ctx, query, userID, phone)
	if err != nil {
		return errors.WrapFail(err, "update user phone")
	}
	if res.RowsAffected() == 0 {
		return user.NotFoundError{
			Err: errors.Errorf("no user rows with %v", errors.Token("id", userID)),
		}
	}

	return nil
}

func (r *repo) GetByEmail(ctx context.Context, email user.Email) (user.User, error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`
	row := r.db.QueryRow(ctx, query, email)

	var u user.User
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.User{}, user.NotFoundError{
			Err: errors.Errorf("no user rows with %v", errors.Token("email", email)),
		}
	case err != nil:
		return user.User{}, errors.WrapFail(err, "fetch user by email")
	}

	return u, nil
}
