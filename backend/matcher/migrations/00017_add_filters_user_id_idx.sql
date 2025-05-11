-- +goose Up
-- +goose StatementBegin
create index if not exists idx_filters_user_id on filters(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_filters_user_id;
-- +goose StatementEnd