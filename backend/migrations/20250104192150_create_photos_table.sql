-- +goose Up
create table if not exists Photos (
    id BIGINT PRIMARY KEY,
    user_id BIGINT,
    url VARCHAR(255) NOT NULL,
    mime_type VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- +goose Down
drop table if exists Photos;
