INSERT INTO tags (
	id,
	name,
	color,
	user_id,
	created_at,
	updated_at,
	deleted_at 
) VALUES ($1, $2, $3, $4, $5, $6, $7)
