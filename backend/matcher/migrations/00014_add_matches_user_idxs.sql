-- +goose Up
-- +goose StatementBegin
create index concurrently if not exists idx_matches_first_user on matches (first_user_id);
create index concurrently if not exists idx_matches_second_user on matches (second_user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index concurrently if exists idx_matches_first_user;
drop index concurrently if exists idx_matches_second_user;
-- +goose StatementEnd
