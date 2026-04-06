package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
)

type NotebookRepository struct {
	db *sql.DB
}

func NewNotebookRepository(d *sql.DB) *NotebookRepository {
	return &NotebookRepository{db: d}
}

//go:embed _query/list_notebook.sql
var listNotebookQuery string

//go:embed _query/create_notebook.sql
var createNotebookQuery string

//go:embed _query/delete_notebook.sql
var deleteNotebookQuery string

//go:embed _query/update_notebook.sql
var updateNotebookQuery string

//go:embed _query/find_by_user_id_notebook.sql
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
