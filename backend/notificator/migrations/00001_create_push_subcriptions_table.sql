-- +goose Up
-- +goose StatementBegin
create table if not exists push_subscriptions
(
    endpoint     text primary key,
    p256dh       text,
    auth_key     text,
    recipient_id uuid
);

create index recipient_id_idx on push_subscriptions (recipient_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists push_subscriptions
-- +goose StatementEnd
