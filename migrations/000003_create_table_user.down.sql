CREATE TABLE IF NOT EXISTS users (
    id integer primary key,
    email varchar(128) unique,
    password varchar(256),
    name varchar(128),
);
