-- +goose Up
-- +goose StatementBegin
INSERT INTO role_ (title) VALUES ('user'), ('admin')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM role_ WHERE title IN ('user', 'admin')
-- +goose StatementEnd
