create table tokens (
    id int auto_increment not null PRIMARY KEY,
    user_id int not null,
    token varchar(256) not null,
    user_type varchar(16) not null
)
