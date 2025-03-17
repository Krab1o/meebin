-- +goose Up
-- +goose StatementBegin
CREATE TABLE session_ (
    id_session SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES user_(id) ON DELETE CASCADE,
    expiration_time TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE session_;
-- +goose StatementEnd
