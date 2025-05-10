-- +goose Up
-- +goose StatementBegin
create table if not exists idempotency_requests
(
    idempotency_key uuid        not null primary key,
    created_at      timestamptz not null default now(),
    result          bytea       null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists idempotency_requests;
-- +goose StatementEnd
