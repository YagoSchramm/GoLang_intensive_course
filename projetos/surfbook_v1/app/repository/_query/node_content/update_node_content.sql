UPDATE nodes_contents
SET
  updated_at = CURRENT_TIMESTAMP,
  content_id = CASE WHEN $3::uuid IS NOT NULL THEN $3::uuid ELSE content_id END,
  notebook_id = CASE WHEN $4::uuid IS NOT NULL THEN $4::uuid ELSE notebook_id END
WHERE id = $1 AND user_id = $2;
