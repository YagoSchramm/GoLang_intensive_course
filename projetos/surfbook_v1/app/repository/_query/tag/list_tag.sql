SELECT 
    id, name, color, user_id, deleted_at, created_at, updated_at
FROM tags
WHERE user_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
