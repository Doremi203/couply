-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id         uuid        not null primary key,
    email      varchar     not null unique,
    password   varchar     not null,
    phone      varchar(16) null unique,
    created_at timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
