package tokenpostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
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

func (r *repo) Get(ctx context.Context, refreshValue token.RefreshValue) (token.Refresh, error) {
	const query = `
		SELECT token, user_id, expires_at
		FROM refresh_tokens
		WHERE token = $1
	`
	row := r.db.QueryRow(ctx, query, refreshValue)

	var refresh token.Refresh
	err := row.Scan(&refresh.Token, &refresh.UserID, &refresh.ExpiresAt)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return token.Refresh{}, token.ErrRefreshTokenNotFound
	case err != nil:
		return token.Refresh{}, errors.WrapFail(err, "fetch refresh token")
	}

	return refresh, nil
}

func (r *repo) Create(ctx context.Context, refresh token.Refresh) error {
	const query = `
		INSERT INTO refresh_tokens (token, user_id, expires_at)
		VALUES ($1, $2, $3)
		ON CONFLICT (token) DO NOTHING
	`
	res, err := r.db.Exec(
		ctx, query,
		refresh.Token,
		refresh.UserID,
		refresh.ExpiresAt,
	)
	if err != nil {
		return errors.WrapFail(err, "create refresh token")
	}
	if res.RowsAffected() == 0 {
		return errors.Error("refresh token already exists")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, refresh token.Refresh) error {
	const query = `
		DELETE FROM refresh_tokens
		WHERE token = $1
	`
	res, err := r.db.Exec(ctx, query, refresh.Token)
	if err != nil {
		return errors.WrapFail(err, "delete refresh token")
	}
	if res.RowsAffected() == 0 {
		return token.ErrRefreshTokenNotFound
	}

	return nil
}
