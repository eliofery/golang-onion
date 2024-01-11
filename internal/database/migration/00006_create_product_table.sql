-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    title       TEXT,
    description TEXT,
    image       TEXT,
    price       NUMERIC(10, 2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
