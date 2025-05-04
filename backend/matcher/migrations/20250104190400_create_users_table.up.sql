create table if not exists users (
    id bigserial primary key,
    name text not null,
    age int check (age >= 18),
    gender int not null,
    location text not null,
    bio text,
    goal int,
    zodiac int,
    height int check (height > 0),
    education int,
    children int,
    alcohol int,
    smoking int,
    hidden boolean not null,
    verified boolean not null,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);
