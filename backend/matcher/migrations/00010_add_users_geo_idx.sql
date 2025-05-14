-- +goose Up
-- +goose StatementBegin
create index if not exists idx_users_geo on users using gist (ll_to_earth(latitude, longitude));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_users_geo;
-- +goose StatementEnd
