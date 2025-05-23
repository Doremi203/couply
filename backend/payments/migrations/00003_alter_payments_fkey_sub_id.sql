-- +goose Up
-- +goose StatementBegin
alter table payments add constraint fk_filters_subscription_id foreign key (subscription_id) references subscriptions(id) on delete cascade;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table payments drop constraint if exists fk_filters_subscription_id;
-- +goose StatementEnd
