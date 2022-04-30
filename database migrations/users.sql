create table users (
    id int auto_increment not null PRIMARY KEY,
    email varchar(64) not null,
    phone varchar(16) not null,
    name varchar(64),
    surname varchar(64),
    photo_id int,
    country varchar(64),
    city varchar(64),
    address varchar(256)
);

ALTER TABLE users
    ADD COLUMN createdTime timestamp;

ALTER TABLE users
    ADD COLUMN password varchar(60) AFTER email;