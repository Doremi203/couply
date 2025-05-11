-- +goose Up
-- +goose StatementBegin
alter table filters add constraint filters_pkey primary key (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table filters drop constraint if exists filters_pkey;
-- +goose StatementEnd
