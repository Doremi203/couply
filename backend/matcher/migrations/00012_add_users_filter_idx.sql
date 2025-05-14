-- +goose Up
-- +goose StatementBegin
create index if not exists idx_users_zodiac on users (zodiac);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_users_zodiac;
-- +goose StatementEnd
