-- +goose Up
-- +goose StatementBegin
create table if not exists payments
(
    id uuid primary key,
    user_id uuid not null,
    subscription_id uuid not null,
    amount bigint not null,
    currency char(3) not null,
    status int not null,
    gateway_id text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists payments;
-- +goose StatementEnd
