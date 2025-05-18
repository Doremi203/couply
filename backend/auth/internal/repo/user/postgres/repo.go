package userpostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
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
		INSERT INTO users (id, email, password, phone)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (email) DO NOTHING;
	`
	res, err := r.db.Exec(
		ctx, query,
		u.ID,
		u.Email,
		u.Password,
		u.Phone,
	)
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
		return errors.Wrapf(
			user.ErrNotFound,
			"no rows with %v",
			errors.Token("id", userID),
		)
	}

	return nil
}

func (r *repo) GetByEmail(ctx context.Context, email user.Email) (user.User, error) {
	query := `
		SELECT id, email, phone, password
		FROM users
		WHERE email = $1
	`
	row := r.db.QueryRow(ctx, query, email)

	var u user.User
	var phone *string
	err := row.Scan(&u.ID, &u.Email, &phone, &u.Password)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.User{}, errors.Wrapf(user.ErrNotFound,
			"no rows with %v",
			errors.Token("email", email),
		)
	case err != nil:
		return user.User{}, errors.WrapFail(err, "fetch user by email")
	}
	if phone != nil {
		u.Phone = user.Phone(*phone)
	}

	return u, nil
}

func (r *repo) GetByPhone(ctx context.Context, phone user.Phone) (user.User, error) {
	query := `
		SELECT id, email, phone, password
		FROM users
		WHERE phone = $1
	`
	row := r.db.QueryRow(ctx, query, phone)

	var u user.User
	var p *string
	err := row.Scan(&u.ID, &u.Email, &p, &u.Password)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.User{}, errors.Wrapf(
			user.ErrNotFound,
			"no rows with %v",
			errors.Token("phone", phone),
		)
	case err != nil:
		return user.User{}, errors.WrapFail(err, "fetch user by email")
	}
	if p != nil {
		u.Phone = user.Phone(*p)
	}

	return u, nil
}

func (r *repo) GetByOAuthProviderUserID(ctx context.Context, provider oauth.Provider, providerUserID oauth.ProviderUserID) (user.User, error) {
	const query = `
		SELECT u.id, u.email, u.phone, u.password
		FROM users u
		JOIN user_oauth_accounts oa ON u.id = oa.user_id
		WHERE oa.provider = $1 AND oa.provider_user_id = $2;
	`
	row := r.db.QueryRow(ctx, query, provider, providerUserID)

	var u user.User
	err := row.Scan(&u.ID, &u.Email, &u.Phone, &u.Password)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.User{}, errors.Wrapf(
			user.ErrNotFound,
			"no rows with %v and %v",
			errors.Token("provider", provider),
			errors.Token("provider_user_id", providerUserID),
		)

	case err != nil:
		return user.User{}, errors.WrapFail(err, "fetch oauth account by provider and provider user id")
	}

	return u, nil
}
