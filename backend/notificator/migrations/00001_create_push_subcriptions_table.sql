-- +goose Up
-- +goose StatementBegin
create table if not exists push_subscriptions
(
    endpoint text primary key,
    p256dh   text,
    auth_key text,
    user_id  uuid
);

create index user_id_idx on push_subscriptions (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists push_subscriptions
-- +goose StatementEnd
