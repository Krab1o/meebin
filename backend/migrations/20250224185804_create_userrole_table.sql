-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_role_ (
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES id_user_(id) ON DELETE CASCADE,
    id_role INTEGER NOT NULL REFERENCES id_role_(id) ON DELETE NO ACTION,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_role_;
-- +goose StatementEnd
