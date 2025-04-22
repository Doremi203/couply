create table if not exists matches (
    main_user_id bigint,
    chosen_user_id bigint,
    approved boolean,
    foreign key (main_user_id) references Users(id) on delete cascade ,
    foreign key (chosen_user_id) references Users(id) on delete cascade
);