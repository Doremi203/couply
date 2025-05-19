-- +goose Up
-- +goose StatementBegin
create table if not exists user_blocks
(
    id uuid primary key,
    blocked_id uuid not null,
    message text not null,
    created_at timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_blocks;
-- +goose StatementEnd