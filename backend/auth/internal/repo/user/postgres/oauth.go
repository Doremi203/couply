package userpostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

func NewOAuthAccountRepo(
	db postgres.Client,
) *oauthAccountRepo {
	return &oauthAccountRepo{
		db: db,
	}
}

type oauthAccountRepo struct {
	db postgres.Client
}

func (r *oauthAccountRepo) Create(ctx context.Context, account user.OAuthAccount) error {
	const query = `
		INSERT INTO user_oauth_accounts (user_id, provider, provider_user_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (provider, provider_user_id) DO NOTHING;
	`
	res, err := r.db.Exec(
		ctx, query,
		account.UserID,
		account.Provider,
		account.ProviderUserID,
	)
	if err != nil {
		return errors.WrapFail(err, "save oauth account")
	}
	if res.RowsAffected() == 0 {
		return errors.Wrapf(
			user.ErrOAuthAccountAlreadyExists,
			"oauth account with %v and %v",
			errors.Token("provider", account.Provider),
			errors.Token("provider_user_id", account.ProviderUserID),
		)
	}

	return nil
}

func (r *oauthAccountRepo) Upsert(ctx context.Context, account user.OAuthAccount) error {
	const query = `
		INSERT INTO user_oauth_accounts (user_id, provider, provider_user_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (provider, provider_user_id) DO UPDATE 
		SET user_id = $1;
	`
	_, err := r.db.Exec(
		ctx, query,
		account.UserID,
		account.Provider,
		account.ProviderUserID,
	)
	if err != nil {
		return errors.WrapFail(err, "upsert oauth account")
	}

	return nil
}

func (r *oauthAccountRepo) GetByProviderUserID(
	ctx context.Context,
	provider oauth.ProviderType,
	providerUserID oauth.ProviderUserID,
) (user.OAuthAccount, error) {
	const query = `
		SELECT provider, provider_user_id
		FROM user_oauth_accounts
		WHERE provider = $1 AND provider_user_id = $2
	`
	row := r.db.QueryRow(ctx, query, provider, providerUserID)

	var account user.OAuthAccount
	err := row.Scan(&account.Provider, &account.ProviderUserID)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.OAuthAccount{}, errors.Wrapf(
			user.ErrOAuthAccountNotFound,
			"with %v and %v",
			errors.Token("provider", provider),
			errors.Token("provider_user_id", providerUserID),
		)

	case err != nil:
		return user.OAuthAccount{}, errors.WrapFail(err, "fetch oauth account by provider and provider user id")
	}

	return account, nil
}

func (r *oauthAccountRepo) GetByUserID(ctx context.Context, userID user.ID) (user.OAuthAccount, error) {
	const query = `
		SELECT provider, provider_user_id
		FROM user_oauth_accounts
		WHERE user_id = $1
	`
	row := r.db.QueryRow(ctx, query, userID)

	var account user.OAuthAccount
	err := row.Scan(&account.Provider, &account.ProviderUserID)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return user.OAuthAccount{}, errors.Wrapf(
			user.ErrOAuthAccountNotFound,
			"with %v",
			errors.Token("user_id", userID),
		)
	case err != nil:
		return user.OAuthAccount{}, errors.WrapFail(err, "fetch oauth account by user id")
	}

	return account, nil
}
