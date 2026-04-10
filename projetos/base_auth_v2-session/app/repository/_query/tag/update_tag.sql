UPDATE tags
SET
  updated_at = CURRENT_TIMESTAMP,
  name = CASE WHEN $3::text IS NOT NULL AND $3::text <> '' THEN $3::text ELSE name END,
  color = CASE WHEN $4::text IS NOT NULL AND $4::text <> '' THEN $4::text ELSE color END
WHERE id = $1 AND user_id = $2;
