-- +goose Up
-- +goose StatementBegin
create table if not exists phone_confirmation_requests
(
    user_id    uuid        not null,
    phone      varchar(16) not null,
    code       varchar     not null,
    created_at timestamptz not null default now(),
    expires_in interval    not null,
    primary key (phone, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists phone_confirmation_requests;
-- +goose StatementEnd
