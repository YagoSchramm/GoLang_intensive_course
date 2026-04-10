package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/base-auth-v2-session/model"
)

type TagRepository struct {
	db *sql.DB
}

func NewTagRepository(d *sql.DB) *TagRepository {
	return &TagRepository{db: d}
}

//go:embed _query/tag/list_tag.sql
var listTagQuery string

//go:embed _query/tag/create_tag.sql
var createTagQuery string

//go:embed _query/tag/delete_tag.sql
var deleteTagQuery string

//go:embed _query/tag/update_tag.sql
var updateTagQuery string

//go:embed _query/tag/find_by_user_id_tag.sql
var findByUserIDAndIDTagQuery string

func (r *TagRepository) Create(ctx context.Context, tag *model.TagEntity) error {
	_, err := r.db.Exec(
		createTagQuery,
		tag.TagID,
		tag.Name,
		tag.Color,
		tag.UserID,
		tag.CreatedAt,
		tag.UpdatedAt,
		tag.DeletedAt,
	)
	return err
}

func (r *TagRepository) Update(ctx context.Context, updateIt model.UpdateTagDTO) error {
	_, err := r.db.Exec(
		updateTagQuery,
		updateIt.TagID,
		updateIt.UserID,
		updateIt.Name,
		updateIt.Color,
	)
	return err
}

func (r *TagRepository) Delete(ctx context.Context, deleteIt model.DeleteTagDTO) error {
	_, err := r.db.Exec(
		deleteTagQuery,
		deleteIt.TagID,
		deleteIt.UserID,
	)
	return err
}

func (r *TagRepository) FindByUserIdTagId(ctx context.Context, input model.FindTagFromUserDTO) (*model.TagEntity, error) {
	rows, err := r.db.QueryContext(ctx, findByUserIDAndIDTagQuery, input.UserID, input.TagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tag model.TagEntity
	for rows.Next() {
		err = rows.Scan(
			&tag.TagID,
			&tag.Name,
			&tag.Color,
			&tag.UserID,
			&tag.DeletedAt,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) ListTags(ctx context.Context, input model.ListTagsFromUserDTO) ([]*model.TagEntity, error) {
	rows, err := r.db.QueryContext(ctx, listTagQuery, input.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tagList []*model.TagEntity
	for rows.Next() {
		var tag model.TagEntity
		err := rows.Scan(
			&tag.TagID,
			&tag.Name,
			&tag.Color,
			&tag.UserID,
			&tag.DeletedAt,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tagList = append(tagList, &tag)
	}
	return tagList, nil
}
