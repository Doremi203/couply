package userpostgres

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepo(
	db *pgxpool.Pool,
) *repo {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *pgxpool.Pool
}

func (p *repo) Create(ctx context.Context, u user.User) error {
	query := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3);
	`
	_, err := p.db.Exec(ctx, query, u.ID, u.Email, u.Password)
	switch {
	case postgres.IsUniqueViolationError(err):
		return errors.Wrapf(user.ErrAlreadyExists, "user with email %s", u.Email)
	case err != nil:
		return errors.Wrap(err, "failed to save user")
	}

	return nil
}

func (p *repo) GetByEmail(ctx context.Context, email user.Email) (user.User, error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`
	row := p.db.QueryRow(ctx, query, email)

	var u user.User
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.User{}, errors.Wrapf(user.ErrNotFound, "with email %s", email)
	case err != nil:
		return user.User{}, errors.Wrap(err, "failed to fetch user by email")
	}

	return u, nil
}
