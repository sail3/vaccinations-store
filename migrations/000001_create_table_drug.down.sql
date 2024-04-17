CREATE TABLE IF NOT EXISTS drug (
    id integer primary key,
    name varchar(128) not null,
    approved varchar (128) not null,
    min_dose integer not null,
    max_dose integer not null,
    available_at timestamp not null
);