package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/base-auth-v2-session/model"
)

type MetaContentRepository struct {
	db *sql.DB
}

func NewMetaContentRepository(d *sql.DB) *MetaContentRepository {
	return &MetaContentRepository{db: d}
}

//go:embed _query/meta_content/list_meta_content.sql
var listMetaContentQuery string

//go:embed _query/meta_content/create_meta_content.sql
var createMetaContentQuery string

//go:embed _query/meta_content/delete_meta_content.sql
var deleteMetaContentQuery string

//go:embed _query/meta_content/update_meta_content.sql
var updateMetaContentQuery string

//go:embed _query/meta_content/find_by_user_id_meta_content.sql
var findByUserIDAndIDMetaContentQuery string

func (r *MetaContentRepository) Create(ctx context.Context, input model.MetaContentEntity) error {
	_, err := r.db.Exec(createMetaContentQuery,
		input.MetaContentID,
		input.NotebookID,
		input.UserID,
		input.Icon,
		input.Name,
		input.CreatedAt,
		input.UpdatedAt,
		input.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r *MetaContentRepository) ListMetaContentById(ctx context.Context, input model.ListMetaContentByIdDTO) (*model.MetaContentEntity, error) {
	rows, err := r.db.QueryContext(
		ctx,
		findByUserIDAndIDMetaContentQuery,
		input.UserID,
		input.MetaContentID,
	)
	if err != nil {
		return nil, err
	}
	var mc model.MetaContentEntity
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&mc.MetaContentID,
			&mc.UserID,
			&mc.Name,
			&mc.NotebookID,
			&mc.Icon,
			&mc.CreatedAt,
			&mc.UpdatedAt,
		)
	}
	if err != nil {
		return nil, err
	}
	return &mc, err
}
func (r *MetaContentRepository) ListMetaContentByUserId(ctx context.Context, input model.ListMetaContentFromUserDTO) (*[]model.MetaContentEntity, error) {
	rows, err := r.db.QueryContext(
		ctx,
		listMetaContentQuery,
		input.UserID,
	)
	if err != nil {
		return nil, err
	}
	var metacontentList []model.MetaContentEntity
	for rows.Next() {
		var metaContent model.MetaContentEntity
		err := rows.Scan(
			&metaContent.MetaContentID,
			&metaContent.UserID,
			&metaContent.Name,
			&metaContent.NotebookID,
			&metaContent.Icon,
			&metaContent.CreatedAt,
			&metaContent.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		metacontentList = append(metacontentList, metaContent)
	}
	return &metacontentList, err
}
func (r *MetaContentRepository) Delete(ctx context.Context, deleteIt model.DeleteMetaContentDTO) error {
	_, err := r.db.QueryContext(
		ctx,
		deleteMetaContentQuery,
		deleteIt.MetaContentID,
		deleteIt.UserID,
	)
	return err
}
func (r *MetaContentRepository) Update(ctx context.Context, updateIt model.UpdateMetaContentDTO) error {
	_, err := r.db.QueryContext(
		ctx,
		updateMetaContentQuery,
		updateIt.MetaContentID,
		updateIt.UserID,
		updateIt.Name,
		updateIt.Icon,
	)
	return err
}
