-- +goose Up
create table if not exists Matches (
    main_user_id BIGINT,
    chosen_user_id BIGINT,
    approved BOOLEAN,
    FOREIGN KEY (main_user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (chosen_user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- +goose Down
drop table if exists Matches;