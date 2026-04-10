SELECT 
   id, content_id, user_id, notebook_id, deleted_at, created_at, updated_at
FROM nodes_contents
WHERE user_id = $1 and id = $2
