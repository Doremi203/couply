create table if not exists interests (
    user_id bigint,
    type text,
    value int,
    foreign key (user_id) references users(id) on delete cascade
);