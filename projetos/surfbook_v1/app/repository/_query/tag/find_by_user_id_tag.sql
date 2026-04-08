SELECT 
   id, name, color, user_id, deleted_at, created_at, updated_at
FROM tags
WHERE user_id = $1 and id = $2
