-- +goose Up
-- +goose StatementBegin
create index concurrently if not exists idx_users_geo on users using gist (ll_to_earth(latitude, longitude));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index concurrently if exists idx_users_geo;
-- +goose StatementEnd
