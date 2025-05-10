-- +goose Up
-- +goose StatementBegin
create extension if not exists cube;
create extension if not exists earthdistance;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop extension if exists earthdistance;
drop extension if exists cude;
-- +goose StatementEnd