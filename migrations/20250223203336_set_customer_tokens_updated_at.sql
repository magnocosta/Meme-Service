-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER set_customer_tokens_updated_at
BEFORE UPDATE ON customer_tokens
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_customer_tokens_updated_at ON customer_tokens;
-- +goose StatementEnd
