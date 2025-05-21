-- +goose Up
-- +goose StatementBegin
alter table user_blocks add column status int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table user_blocks drop column status;
-- +goose StatementEnd