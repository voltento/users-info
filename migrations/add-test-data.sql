DELETE FROM users;

INSERT INTO users (first_name, last_name, email, country_code)
VALUES ('John1', 'Mazler1', 'john1@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;

INSERT INTO users (first_name, last_name, email, country_code)
VALUES ('John2', 'Mazler2', 'john2@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;

INSERT INTO users (first_name, last_name, email, country_code)
VALUES ('John3', 'Mazler3', 'john3@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;

INSERT INTO users (first_name, last_name, email, country_code)
VALUES ('John3', 'Mazler3', 'john3@test.ru', 'EN') ON CONFLICT (email) DO NOTHING;