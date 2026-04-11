package service_test

import (
	"database/sql"

	"github.com/google/uuid"
)

func createUser(db *sql.DB, id uuid.UUID, name, email, password string) error {
	query := `INSERT INTO users (id, name, email, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := db.Exec(query, id, name, email, nil)
	return err
}
