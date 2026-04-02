package repository

import (
	"context"
	"database/sql"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
)

type NotebookRepository struct {
	db *sql.DB
}

func NewNotebookRepository(d *sql.DB) *NotebookRepository {
	return &NotebookRepository{db: d}
}

var listNotebookQuery string

var createNotebookQuery string

var deleteNotebookQuery string

var updateNotebookQuery string

var findByUserIDAndIDNotebookQuery string

func (r *NotebookRepository) Create(ctx context.Context, notebook *model.NotebookEntity) error {
	_, err := r.db.Exec(
		createNotebookQuery,
		notebook.NotebookID,
		notebook.UserID,
		notebook.Icon,
		notebook.Name,
		notebook.Description,
		notebook.Image,
		notebook.CreatedAt,
		notebook.UpdatedAt,
		notebook.DeletedAt,
	)
	return err
}
