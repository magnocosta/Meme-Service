INSERT 
  INTO customers(name, email)
  VALUES($1, $2)
  RETURNING id, external_id, name, email, created_at
