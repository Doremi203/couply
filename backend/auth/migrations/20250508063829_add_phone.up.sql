alter table users add column phone varchar(16) not null default '';
create unique index users_phone_uidx on users(phone);