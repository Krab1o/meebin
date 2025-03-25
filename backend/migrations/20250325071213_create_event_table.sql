-- +goose Up
-- +goose StatementBegin
CREATE TABLE event_ (
    id SERIAL PRIMARY KEY,
    status_ INTEGER NOT NULL REFERENCES event_status_(id) ON DELETE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event_;
-- +goose StatementEnd
