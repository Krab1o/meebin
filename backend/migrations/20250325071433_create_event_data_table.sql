-- +goose Up
-- +goose StatementBegin
CREATE TABLE event_data_ (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES event_(id) ON DELETE CASCADE,
    latitude FLOAT NOT NULL,
    longtitude FLOAT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description_ TEXT,
    time_called TIMESTAMP,
    time_utilized TIMESTAMP,
    caller_id INTEGER REFERENCES user_(id) ON DELETE NO ACTION,
    utilizator_id INTEGER REFERENCES user_(id) ON DELETE NO ACTION
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event_data_;
-- +goose StatementEnd
