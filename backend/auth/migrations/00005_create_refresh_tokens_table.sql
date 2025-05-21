-- +goose Up
-- +goose StatementBegin
create table refresh_tokens (
    token varchar(255) not null primary key,
    user_id uuid not null,
    expires_at timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table refresh_tokens;
-- +goose StatementEnd
