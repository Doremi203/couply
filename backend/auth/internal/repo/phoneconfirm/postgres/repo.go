package phoneconfirmpostgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
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

func (r *repo) UpsertRequest(ctx context.Context, request phoneconfirm.Request) error {
	const query = `
		INSERT INTO phone_confirmation_requests (user_id, phone, code, expires_in)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, phone)
		DO UPDATE SET
		  code = EXCLUDED.code,
		  expires_in = EXCLUDED.expires_in,
		  created_at = NOW()
	`
	_, err := r.db.Exec(ctx, query, request.UserID, request.Phone, request.HashedCode, request.Code.ExpiresIn)
	if err != nil {
		return errors.WrapFail(err, "upsert phone confirmation request")
	}
	return nil
}

func (r *repo) GetRequest(ctx context.Context, userID user.ID, phone user.Phone) (phoneconfirm.Request, error) {
	query := `
		SELECT 
		    user_id,
		    phone,
		    code,
		    expires_in,
		    created_at
		FROM phone_confirmation_requests
		WHERE user_id = $1 AND phone = $2
	`
	row := r.db.QueryRow(ctx, query, userID, phone)

	var req phoneconfirm.Request
	err := row.Scan(
		&req.UserID,
		&req.Phone,
		&req.HashedCode,
		&req.Code.ExpiresIn,
		&req.CreatedAt,
	)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return phoneconfirm.Request{}, phoneconfirm.ErrRequestNotFound
	case err != nil:
		return phoneconfirm.Request{}, errors.WrapFail(err, "fetch user by email")
	}

	return req, nil
}

func (r *repo) DeleteRequest(ctx context.Context, userID user.ID, phone user.Phone) error {
	const query = `
		DELETE FROM phone_confirmation_requests
		WHERE user_id = $1 AND phone = $2
	`
	res, err := r.db.Exec(ctx, query, userID, phone)
	if err != nil {
		return errors.WrapFail(err, "delete phone confirmation request")
	}
	if res.RowsAffected() == 0 {
		return phoneconfirm.ErrRequestNotFound
	}

	return nil
}
