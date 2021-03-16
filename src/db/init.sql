CREATE DATABASE stat_db;
CREATE USER stat_dba WITH SUPERUSER NOINHERIT ENCRYPTED PASSWORD 'avito_tradex';
GRANT ALL PRIVILEGES ON DATABASE stat_db TO stat_dba;


CREATE TABLE statistics(
    id SERIAL PRIMARY KEY,
    date date NOT NULL,
    views int CONSTRAINT views_validation CHECK (views >= 0),
    clicks int CONSTRAINT clicks_validation CHECK (clicks >= 0),
    cost money CONSTRAINT cost_validation CHECK (cost::numeric >= 0)
);