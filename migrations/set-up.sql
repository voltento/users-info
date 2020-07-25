CREATE TABLE IF NOT EXISTS users_info (
    userid    serial PRIMARY KEY,
    firstname varchar(40) NOT NULL,
    lastname varchar(40) NOT NULL,
    email varchar(40) NOT NULL UNIQUE,
    countrycode varchar(5) NOT NULL
);