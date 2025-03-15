package user

import (
	"context"
	user2 "github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/repo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresRepo(
	db *pgxpool.Pool,
) *postgresRepo {
	return &postgresRepo{
		db: db,
	}
}

type postgresRepo struct {
	db *pgxpool.Pool
}

func (p *postgresRepo) Save(ctx context.Context, u user2.User) error {
	query := `
		INSERT INTO users (uid, email, password)
		VALUES ($1, $2, $3)
	`
	_, err := p.db.Exec(ctx, query, u.UID, u.Email, u.Password)
	var pgErr pgconn.PgError
	switch {
	case errors.As(err, &pgErr) && pgErr.Code == "23505":
		return errors.Wrapf(repo.ErrAlreadyExists, "user with email %s", u.Email)
	case err != nil:
		return errors.Wrap(err, "failed to save user")
	}

	return nil
}

func (p *postgresRepo) GetByEmail(ctx context.Context, email user2.Email) (user2.User, error) {
	query := `
		SELECT uid, email, password
		FROM users
		WHERE email = $1
	`
	row := p.db.QueryRow(ctx, query, email)

	var u user2.User
	err := row.Scan(&u.UID, &u.Email, &u.Password)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user2.User{}, errors.Wrapf(repo.ErrNotFound, "user with email %s", email)
	case err != nil:
		return user2.User{}, errors.Wrap(err, "failed to fetch user by email")
	}

	return u, nil
}
