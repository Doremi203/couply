-- +goose Up
-- +goose StatementBegin
create table if not exists filters
(
    user_id         uuid ,
    gender_priority int,
    min_age         int,
    max_age         int,
    min_height      int,
    max_height      int,
    min_distance_km int,
    max_distance_km int,
    goal            int,
    zodiac          int,
    education       int,
    children        int,
    alcohol         int,
    smoking         int,
    only_verified   boolean,
    only_premium    boolean,
    created_at      timestamptz default current_timestamp,
    updated_at      timestamptz default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists filters;
-- +goose StatementEnd
