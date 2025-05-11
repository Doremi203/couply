-- +goose Up
-- +goose StatementBegin
create index concurrently if not exists idx_users_zodiac on users (zodiac);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index concurrently if exists idx_users_zodiac;
-- +goose StatementEnd
