create table if not exists photos (
    user_id uuid,
    order_number bigint,
    url text not null,
    mime_type text not null,
    uploaded_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    primary key (user_id, order_number),
    foreign key (user_id) references Users(id) on delete cascade
);
