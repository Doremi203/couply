-- +goose Up
-- +goose StatementBegin
create table if not exists filter_interests
(
    user_id uuid,
    type    text,
    value   int,
    foreign key (user_id) references users (id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists filter_interests;
-- +goose StatementEnd
