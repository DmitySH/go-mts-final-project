-- +goose Up
-- +goose StatementBegin
INSERT INTO auto (name)
VALUES ('Солярис'),
       ('Лада'),
       ('Беха пятерка'),
       ('Волк вагон'),
       ('Китаец');

INSERT INTO driver
VALUES (gen_random_uuid(), 'Петров Илья Никитич', (SELECT id FROM auto WHERE name = 'Солярис'), 20, 20),
       (gen_random_uuid(), 'Фомина Виктория Артёмовна', (SELECT id FROM auto WHERE name = 'Волк вагон'), 10, 20),
       (gen_random_uuid(), 'Абрамов Иван Степанович', (SELECT id FROM auto WHERE name = 'Китаец'), 20, 10),
       (gen_random_uuid(), 'Скворцова Алиса Дмитриевна', (SELECT id FROM auto WHERE name = 'Китаец'), 12.32, 32.12),
       (gen_random_uuid(), 'Кузнецов Лев Дмитриевич', (SELECT id FROM auto WHERE name = 'Беха пятерка'), 2, 80),
       (gen_random_uuid(), 'Зуев Глеб Александрович', (SELECT id FROM auto WHERE name = 'Лада'), 3.4124, 2.321),
       (gen_random_uuid(), 'Гордеева Полина Давидовна', (SELECT id FROM auto WHERE name = 'Лада'), 1, 1),
       (gen_random_uuid(), 'Тарасов Степан Тимофеевич', (SELECT id FROM auto WHERE name = 'Волк вагон'), 65.3, 72);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE driver;
TRUNCATE TABLE auto;

-- +goose StatementEnd
