-- +goose Up
create table if not exists Interests (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    type VARCHAR(255),
    value INT,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- +goose Down
drop table if exists Interests;