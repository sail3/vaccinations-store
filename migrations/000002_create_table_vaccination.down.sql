CREATE TABLE IF NOT EXISTS vaccination (
    id integer primary key,
    name varchar(128) not null,
    drug_id integer not null,
    dose integer not null,
    date datetime not null,
    CONSTRAINT fk_drug
    FOREIGN KEY(drug_id) 
    REFERENCES drug(id)
);