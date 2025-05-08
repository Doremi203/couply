create table if not exists filter_interests (
    user_id uuid,
    type text,
    value int,
    foreign key (user_id) references users(id) on delete cascade
);