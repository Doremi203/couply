create table if not exists likes (
    sender_id uuid not null,
    receiver_id uuid not null,
    message text not null,
    status int not null,
    created_at timestamptz not null default now(),
    foreign key (sender_id) references users(id) on delete cascade,
    foreign key (receiver_id) references users(id) on delete cascade,
    unique(sender_id, receiver_id)
);