package telegram

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/telegram"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/postgres"
	"github.com/google/uuid"
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

var ErrTelegramDataNotFound = errors.Error("telegram data not found")

func (r *repo) UpsertTelegramData(
	ctx context.Context,
	data telegram.Data,
) error {
	const query = `
        INSERT INTO telegram_data (user_id, telegram_id, telegram_username)
        VALUES ($1, $2, $3)
        ON CONFLICT (user_id)
        DO UPDATE SET
            telegram_id = EXCLUDED.telegram_id,
            telegram_username = EXCLUDED.telegram_username
    `
	_, err := r.db.Exec(ctx, query, data.UserID, data.TelegramID, data.TelegramUsername)
	if err != nil {
		return errors.WrapFail(err, "upsert telegram data")
	}
	return nil
}

func (r *repo) GetTelegramData(
	ctx context.Context,
	userID uuid.UUID,
) (telegram.Data, error) {
	const query = `
        SELECT telegram_id, telegram_username
        FROM telegram_data
        WHERE user_id = $1
    `
	row := r.db.QueryRow(ctx, query, userID)

	var data telegram.Data
	err := row.Scan(&data.TelegramID, &data.TelegramUsername)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return telegram.Data{}, ErrTelegramDataNotFound
	case err != nil:
		return telegram.Data{}, errors.WrapFail(err, "get telegram data")
	}

	return data, nil
}
