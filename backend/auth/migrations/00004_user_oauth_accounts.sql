-- +goose Up
-- +goose StatementBegin
alter table users
    alter column password drop not null;

create table user_oauth_accounts
(
    user_id          uuid        not null
        references users (id)
            on delete cascade,
    provider         varchar(50) not null,
    provider_user_id text        not null,
    created_at       timestamptz not null default now(),
    primary key (provider, provider_user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user_oauth_accounts;

-- Update NULL passwords with a placeholder before adding NOT NULL constraint
update users
set password = 'OAUTH_USER_PLACEHOLDER_PASSWORD'
where password is null;

alter table users
    alter column password set not null;
-- +goose StatementEnd
