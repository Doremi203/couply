-- +goose Up
-- +goose StatementBegin
create table if not exists user_views (
    viewer_id uuid not null,
    viewed_id uuid not null,
    viewed_at timestamptz not null default now(),
    primary key (viewer_id, viewed_id),
    foreign key (viewer_id) references users(id) on delete cascade ,
    foreign key (viewed_id) references users(id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_views;
-- +goose StatementEnd
