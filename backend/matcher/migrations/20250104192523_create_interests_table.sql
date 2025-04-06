-- +goose Up
create table if not exists Interests (
    user_id BIGINT,
    type TEXT,
    value TEXT,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- +goose Down
drop table if exists Interests;