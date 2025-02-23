UPDATE customer_tokens 
SET tokens = tokens + $2
WHERE customer_id = $1 
RETURNING customer_id, tokens, updated_at;
