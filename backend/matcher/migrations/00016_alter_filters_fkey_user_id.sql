-- +goose Up
-- +goose StatementBegin
alter table filters add constraint fk_filters_user foreign key (user_id) references users(id) on delete cascade;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table filters drop constraint if exists fk_filters_user;
-- +goose StatementEnd
