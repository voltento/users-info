CREATE TABLE IF NOT EXISTS users (
    user_id    serial PRIMARY KEY,
    first_name varchar(40) NOT NULL,
    last_name varchar(40) NOT NULL,
    email varchar(40) NOT NULL UNIQUE,
    country_code varchar(5) NOT NULL
);