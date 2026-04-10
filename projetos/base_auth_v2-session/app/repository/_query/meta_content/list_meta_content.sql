SELECT 
    id, user_id, name, notebook_id, icon, created_at, updated_at
FROM meta_contents
WHERE user_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC