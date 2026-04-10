package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/google/uuid"
)

type NotebookRepository struct {
	db *sql.DB
}

func NewNotebookRepository(d *sql.DB) *NotebookRepository {
	return &NotebookRepository{db: d}
}

//go:embed _query/notebook/list_notebook.sql
var listNotebookQuery string

//go:embed _query/notebook/create_notebook.sql
var createNotebookQuery string

//go:embed _query/notebook/delete_notebook.sql
var deleteNotebookQuery string

//go:embed _query/notebook/update_notebook.sql
var updateNotebookQuery string

//go:embed _query/notebook/find_by_user_id_notebook.sql
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
func (r *NotebookRepository) Update(ctx context.Context, updateIt model.UpdateNotebookDTO) error {
	_, err := r.db.Exec(
		updateNotebookQuery,
		updateIt.NotebookID,
		updateIt.UserID,
		updateIt.Name,
		updateIt.Description,
		updateIt.Icon,
		updateIt.Image,
	)

	return err
}
func (r *NotebookRepository) Delete(ctx context.Context, deleteIt model.DeleteNotebookDTO) error {
	_, err := r.db.Exec(
		deleteNotebookQuery,
		deleteIt.NotebookID,
		deleteIt.UserID,
	)
	return err
}
func (r *NotebookRepository) FindByUserIdNotebookId(ctx context.Context, notebook_id uuid.UUID, user_id uuid.UUID) (*model.NotebookEntity, error) {
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
