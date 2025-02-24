-- +goose Up
-- +goose StatementBegin
CREATE TABLE session_ (
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES user_(id) ON DELETE CASCADE,
    refresh_token VARCHAR(255) NOT NULL,
    expiration_date TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE session_;
-- +goose StatementEnd
