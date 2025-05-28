-- +goose Up
-- +goose StatementBegin
create table telegram_data (
    user_id uuid primary key not null,
    telegram_id text not null,
    telegram_username text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table telegram_data;
-- +goose StatementEnd
