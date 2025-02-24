-- +goose Up
-- +goose StatementBegin
CREATE TABLE stats_ (
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES user_(id) ON DELETE CASCADE,
    utilize_counter INTEGER NOT NULL,
    report_counter INTEGER NOT NULL,
    rating FLOAT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stats_;
-- +goose StatementEnd
