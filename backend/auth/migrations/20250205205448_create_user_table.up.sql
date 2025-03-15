create table users
(
    uid      uuid primary key,
    email    varchar not null unique,
    password varchar not null
);