package repository

import (
	"context"
	"database/sql"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
	"github.com/google/uuid"
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
func (r *NotebookRepository) Delete(ctx context.Context, notebook_id string, user_id string) error {
	_, err := r.db.Exec(
		deleteNotebookQuery,
		notebook_id,
		user_id,
	)
	return err
}
func (r *NotebookRepository) findByUserIdNotenookId(ctx context.Context, notebook_id string, user_id string) (*model.NotebookEntity, error) {
	var NotebookList []model.NotebookEntity
	rows, err := r.db.QueryContext(ctx, findByUserIDAndIDNotebookQuery, user_id, notebook_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var nb model.NotebookEntity
		err := rows.Scan(
			&nb.NotebookID,
			&nb.UserID,
			&nb.Name,
			&nb.Description,
			&nb.Icon,
			&nb.Image,
			&nb.CreatedAt,
			&nb.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		NotebookList = append(NotebookList, nb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &NotebookList[0], nil
}
func (r *NotebookRepository) ListNotebooks(ctx context.Context, user_id uuid.UUID) ([]*model.NotebookEntity, error) {
	rows, err := r.db.QueryContext(ctx, listNotebookQuery, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var nbList []*model.NotebookEntity
	for rows.Next() {
		var nb model.NotebookEntity
		err := rows.Scan(
			&nb.NotebookID,
			&nb.UserID,
			&nb.Name,
			&nb.Description,
			&nb.Icon,
			&nb.Image,
			&nb.CreatedAt,
			&nb.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		nbList = append(nbList, &nb)
	}
	return nbList, nil
}
