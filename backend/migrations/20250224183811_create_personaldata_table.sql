-- +goose Up
-- +goose StatementBegin
CREATE TABLE personal_data_ (
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES user_(id) ON DELETE CASCADE,
    given_name VARCHAR(127),
    surname VARCHAR(127),
    patronymic VARCHAR(127),
    city VARCHAR(127),
    birthdate DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE personal_data_;
-- +goose StatementEnd