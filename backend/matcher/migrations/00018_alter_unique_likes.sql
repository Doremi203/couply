-- +goose Up
-- +goose StatementBegin
alter table likes drop constraint likes_sender_id_receiver_id_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table likes add constraint likes_sender_id_receiver_id_key unique (sender_id, receiver_id);
-- +goose StatementEnd