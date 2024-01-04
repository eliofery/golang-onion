-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tests;
-- +goose StatementEnd
