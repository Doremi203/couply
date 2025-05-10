-- +goose Up
-- +goose StatementBegin
create table if not exists photos
(
    user_id      uuid,
    order_number bigint,
    object_key   text not null,
    mime_type    text not null,
    uploaded_at  timestamptz default now(),
    primary key (user_id, order_number),
    foreign key (user_id) references users (id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists photos;
-- +goose StatementEnd
