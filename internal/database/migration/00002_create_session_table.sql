-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions
(
    id    SERIAL PRIMARY KEY,
    token VARCHAR(255) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
