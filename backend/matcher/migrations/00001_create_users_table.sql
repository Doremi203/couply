-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id          uuid primary key,
    name        text    not null,
    age         int check (age >= 18),
    gender      int     not null,
    latitude    double precision not null check (latitude >= -90 AND latitude <= 90),
    longitude   double precision not null check (longitude >= -180 AND longitude <= 180),
    bio         text,
    goal        int,
    zodiac      int,
    height      int check (height > 0),
    education   int,
    children    int,
    alcohol     int,
    smoking     int,
    is_hidden   boolean not null,
    is_verified boolean not null,
    is_premium  boolean not null,
    is_blocked  boolean not null,
    created_at  timestamptz default current_timestamp,
    updated_at  timestamptz default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
