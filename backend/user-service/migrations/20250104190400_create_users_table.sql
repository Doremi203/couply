-- +goose Up
create table if not exists Users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT CHECK (age > 18),
    gender INT NOT NULL,
    location VARCHAR(255) NOT NULL,
    bio TEXT,
    goal INT,
    zodiac INT,
    height INT CHECK (height > 0),
    education INT,
    children INT,
    alcohol INT,
    smoking INT,
    hidden BOOLEAN NOT NULL,
    verified BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
drop table if exists Users;
