package userpostgres

import (
	"context"
	"fmt"
	"strings"

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
	var phone *string
	if u.Phone != "" {
		phoneStr := string(u.Phone)
		phone = &phoneStr
	}

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
		phone,
	)
	if err != nil {
		return errors.WrapFail(err, "save user")
	}
	if res.RowsAffected() == 0 {
		return errors.Wrapf(user.ErrAlreadyExists, "user with %v", errors.Token("email", u.Email))
	}

	return nil
}

func (r *repo) Upsert(ctx context.Context, u user.User) error {
	var phone *string
	if u.Phone != "" {
		phoneStr := string(u.Phone)
		phone = &phoneStr
	}

	const query = `
		INSERT INTO users (id, email, password, phone)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE 
		SET email = $2, password = $3, phone = $4;
	`
	_, err := r.db.Exec(
		ctx, query,
		u.ID,
		u.Email,
		u.Password,
		phone,
	)
	if err != nil {
		return errors.WrapFail(err, "upsert user")
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

func (r *repo) GetByAny(ctx context.Context, params user.GetByAnyParams) (user.User, error) {
	if params.AllEmpty() {
		return user.User{}, errors.Error("all get by any params are empty")
	}

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
		SELECT u.id, u.email, u.phone, u.password
		FROM users u
	`)
	args := make([]interface{}, 0, 4)
	argPos := 1

	if params.OAuthUserAccount != nil {
		queryBuilder.WriteString(`
			LEFT JOIN user_oauth_accounts oa ON u.id = oa.user_id
		`)
	}

	queryBuilder.WriteString("WHERE ")
	conditions := make([]string, 0, 4)

	if params.ID != nil {
		conditions = append(conditions, fmt.Sprintf("u.id = $%d", argPos))
		args = append(args, *params.ID)
		argPos++
	}

	if params.Email != "" {
		conditions = append(conditions, fmt.Sprintf("u.email = $%d", argPos))
		args = append(args, params.Email)
		argPos++
	}

	if params.Phone != "" {
		conditions = append(conditions, fmt.Sprintf("u.phone = $%d", argPos))
		args = append(args, params.Phone)
		argPos++
	}

	if params.OAuthUserAccount != nil {
		conditions = append(conditions, fmt.Sprintf("oa.provider = $%d AND oa.provider_user_id = $%d",
			argPos, argPos+1))
		args = append(args, params.OAuthUserAccount.Provider, params.OAuthUserAccount.ProviderUserID)
	}

	queryBuilder.WriteString(strings.Join(conditions, " OR "))

	row := r.db.QueryRow(ctx, queryBuilder.String(), args...)

	var u user.User
	var phonePtr *string
	err := row.Scan(&u.ID, &u.Email, &phonePtr, &u.Password)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.User{}, errors.Wrap(
			user.ErrNotFound,
			"no user found with provided parameters",
		)
	case err != nil:
		return user.User{}, errors.WrapFail(err, "fetch user by parameters")
	}

	if phonePtr != nil {
		u.Phone = user.Phone(*phonePtr)
	}

	return u, nil
}
