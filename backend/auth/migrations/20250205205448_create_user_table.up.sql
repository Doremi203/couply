create table users
(
    id       uuid    not null primary key,
    email    varchar not null unique,
    password varchar not null,
    created_at timestamptz not null default now()
);