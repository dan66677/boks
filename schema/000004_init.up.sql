-- migrations/0001_create_fights_table.up.sql
CREATE TABLE fights (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    fighter1 VARCHAR(100) NOT NULL,
    fighter2 VARCHAR(100) NOT NULL,
    winner VARCHAR(100)
);