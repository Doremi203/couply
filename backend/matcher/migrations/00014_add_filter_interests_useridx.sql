-- +goose Up
-- +goose StatementBegin
create index if not exists idx_filter_interests_user on filter_interests (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_filter_interests_user;
-- +goose StatementEnd
