create table if not exists filters (
    user_id bigint,
    gender_priority int,
    min_age int,
    max_age int,
    min_height int,
    max_height int,
    distance int,
    goal int,
    zodiac int,
    education int,
    children int,
    alcohol int,
    smoking int,
    only_verified boolean,
    only_premium boolean,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);