-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer_tokens (
    customer_id INT PRIMARY KEY REFERENCES customers(id) ON DELETE CASCADE,
    tokens INT NOT NULL DEFAULT 0,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer_tokens;
-- +goose StatementEnd
