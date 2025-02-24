SELECT tokens, updated_at
FROM customer_tokens
WHERE customer_id=$1;
