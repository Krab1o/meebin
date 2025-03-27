-- +goose Up
-- +goose StatementBegin
CREATE TABLE event_status_ (
    id SERIAL PRIMARY KEY,
    title VARCHAR(63)
);

INSERT INTO event_status_ (title) VALUES 
    ('On Moderation'), 
    ('Opened'), 
    ('On Confirmation'), 
    ('Closed');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event_status_
-- +goose StatementEnd
