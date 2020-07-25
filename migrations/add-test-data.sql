DELETE FROM users_info;

INSERT INTO users_info (firstname, lastname, email, countrycode)
VALUES ('John1', 'Mazler1', 'john1@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;

INSERT INTO users_info (firstname, lastname, email, countrycode)
VALUES ('John2', 'Mazler2', 'john2@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;

INSERT INTO users_info (firstname, lastname, email, countrycode)
VALUES ('John3', 'Mazler3', 'john3@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;

INSERT INTO users_info (firstname, lastname, email, countrycode)
VALUES ('John3', 'Mazler3', 'john3@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;