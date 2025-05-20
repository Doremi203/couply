-- +goose Up
-- +goose StatementBegin
create table if not exists subscriptions
(
    id uuid primary key,
    user_id uuid not null,
    plan int not null,
    status int not null,
    auto_renew boolean not null,
    start_date timestamptz not null,
    end_date timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists subscriptions;
-- +goose StatementEnd
