-- +goose Up
create table if not exists Photos (
    user_id BIGINT,
    order_number BIGINT,
    url TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, order_number),
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- +goose Down
drop table if exists Photos;
