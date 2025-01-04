-- +goose Up
create table if not exists Interests (
    id BIGINT PRIMARY KEY,
    user_id BIGINT,
    type INT,
    value INT,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- +goose Down
drop table if exists Interests;