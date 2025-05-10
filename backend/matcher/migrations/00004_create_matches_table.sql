-- +goose Up
-- +goose StatementBegin
create table if not exists matches
(
    first_user_id  uuid        not null,
    second_user_id uuid        not null,
    created_at     timestamptz not null default now(),
    foreign key (first_user_id) references users (id) on delete cascade,
    foreign key (second_user_id) references users (id) on delete cascade,
    unique (first_user_id, second_user_id),
    unique (second_user_id, first_user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists matches;
-- +goose StatementEnd
