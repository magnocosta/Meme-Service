UPDATE customer_tokens 
SET tokens = tokens + $2
WHERE customer_id = $1 
RETURNING tokens, updated_at;
