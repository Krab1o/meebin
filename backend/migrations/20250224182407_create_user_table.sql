-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_ (
    id SERIAL PRIMARY KEY,
	username VARCHAR(64) UNIQUE NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_;
-- +goose StatementEnd