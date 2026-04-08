-- update deleted at
UPDATE nodes_contents
SET 
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND user_id = $2;
