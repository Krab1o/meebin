-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_ (
    id SERIAL PRIMARY KEY,
    title VARCHAR(63)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE role_;
-- +goose StatementEnd
