alter table users add column phone varchar(16) null;
create unique index users_phone_uidx on users(phone);