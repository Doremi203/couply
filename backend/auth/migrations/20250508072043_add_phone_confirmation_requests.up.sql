create table phone_confirmation_requests
(
    user_id    uuid        not null,
    phone      varchar(16) not null,
    code       varchar     not null,
    created_at timestamptz not null default now(),
    expires_in interval    not null,
    primary key (user_id, phone)
);