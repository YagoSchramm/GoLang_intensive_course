package repository

import "database/sql"

type NotebookRepository struct {
	db *sql.DB
}

func NewNotebookRepository(d *sql.DB) *NotebookRepository {
	return &NotebookRepository{db: d}
}
