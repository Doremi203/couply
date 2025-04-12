-- +goose Up
-- +goose StatementBegin
create table users
(
    id       uuid    not null primary key,
    email    varchar not null unique,
    password varchar not null,
    created_at timestamptz not null default now()
);

create table idempotency_requests
(
    idempotency_key uuid not null primary key,
    created_at timestamptz not null default now(),
    result bytea null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
drop table idempotency_requests;
-- +goose StatementEnd
