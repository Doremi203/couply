-- +goose Up
-- +goose StatementBegin
create index if not exists idx_interests_type_value_user on interests (type, value, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_interests_type_value_user;
-- +goose StatementEnd
