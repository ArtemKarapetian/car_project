-- +goose Up
-- +goose StatementBegin
CREATE TABLE people (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        surname VARCHAR(255) NOT NULL,
                        patronymic VARCHAR(255)
);
CREATE TABLE cars (
                      reg_num VARCHAR(255) PRIMARY KEY,
                      mark VARCHAR(255) NOT NULL,
                      model VARCHAR(255) NOT NULL,
                      year INT,
                      owner_id INT REFERENCES people(id),
                      created_at TIMESTAMP,
                      updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars;
DROP TABLE IF EXISTS people;
-- +goose StatementEnd