-- +goose Up
-- +goose StatementBegin
create table if not exists user_block_reasons
(
    block_id uuid,
    reason int not null,
    foreign key (block_id) references user_blocks(id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_block_reasons;
-- +goose StatementEnd