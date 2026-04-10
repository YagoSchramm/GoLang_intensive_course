UPDATE meta_contents
SET
  updated_at = CURRENT_TIMESTAMP,
  name = CASE WHEN $3::text IS NOT NULL AND $3::text <> '' THEN $3::text ELSE name END,
  icon = CASE WHEN $4::text IS NOT NULL AND $5::text <> '' THEN $4::text ELSE icon END,
WHERE id = $1 AND user_id = $2;