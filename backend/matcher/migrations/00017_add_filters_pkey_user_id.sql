-- +goose Up
-- +goose StatementBegin
alter table filters add primary key (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table filters drop constraint filters_pkey;
-- +goose StatementEnd