create table if not exists interests (
    user_id uuid,
    type text,
    value text,
    foreign key (user_id) references users(id) on delete cascade
);