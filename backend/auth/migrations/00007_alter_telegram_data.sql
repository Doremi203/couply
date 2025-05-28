-- +goose Up
-- +goose StatementBegin
ALTER TABLE telegram_data
    ALTER COLUMN telegram_id TYPE bigint USING telegram_id::bigint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE telegram_data
    ALTER COLUMN telegram_id TYPE text USING telegram_id::text;
-- +goose StatementEnd