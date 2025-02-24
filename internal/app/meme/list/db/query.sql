UPDATE customer_tokens 
SET tokens = GREATEST(0, tokens - 1)
WHERE customer_id = $1 
RETURNING tokens;
