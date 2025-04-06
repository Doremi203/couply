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
)