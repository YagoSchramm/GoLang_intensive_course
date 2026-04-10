UPDATE users
SET email = $2,
    password = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;