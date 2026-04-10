package repository

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/base-auth-v2-session/model"
)

type NodeContentRepository struct {
	db *sql.DB
}

func NewNodeContentRepository(d *sql.DB) *NodeContentRepository {
	return &NodeContentRepository{db: d}
}

//go:embed _query/node_content/list_node_content.sql
var listNodeContentQuery string

//go:embed _query/node_content/create_node_content.sql
var createNodeContentQuery string

//go:embed _query/node_content/delete_node_content.sql
var deleteNodeContentQuery string

//go:embed _query/node_content/update_node_content.sql
var updateNodeContentQuery string

//go:embed _query/node_content/find_by_user_id_node_content.sql
var findByUserIDAndIDNodeContentQuery string

func (r *NodeContentRepository) Create(ctx context.Context, node *model.NodeContentEntity) error {
	_, err := r.db.Exec(
		createNodeContentQuery,
		node.NodeID,
		node.ContentID,
		node.UserID,
		node.NotebookID,
		node.CreatedAt,
		node.UpdatedAt,
		node.DeletedAt,
	)
	return err
}

func (r *NodeContentRepository) Update(ctx context.Context, updateIt model.UpdateNodeContentDTO) error {
	_, err := r.db.Exec(
		updateNodeContentQuery,
		updateIt.NodeID,
		updateIt.UserID,
		updateIt.ContentID,
		updateIt.NotebookID,
	)
	return err
}

func (r *NodeContentRepository) Delete(ctx context.Context, deleteIt model.DeleteNodeContentDTO) error {
	_, err := r.db.Exec(
		deleteNodeContentQuery,
		deleteIt.NodeID,
		deleteIt.UserID,
	)
	return err
}

func (r *NodeContentRepository) FindByUserIdNodeId(ctx context.Context, input model.FindNodeContentFromUserDTO) (*model.NodeContentEntity, error) {
	rows, err := r.db.QueryContext(ctx, findByUserIDAndIDNodeContentQuery, input.UserID, input.NodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var node model.NodeContentEntity
	for rows.Next() {
		err = rows.Scan(
			&node.NodeID,
			&node.ContentID,
			&node.UserID,
			&node.NotebookID,
			&node.DeletedAt,
			&node.CreatedAt,
			&node.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &node, nil
}

func (r *NodeContentRepository) ListNodeContents(ctx context.Context, input model.ListNodeContentFromUserDTO) ([]*model.NodeContentEntity, error) {
	rows, err := r.db.QueryContext(ctx, listNodeContentQuery, input.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodeList []*model.NodeContentEntity
	for rows.Next() {
		var node model.NodeContentEntity
		err := rows.Scan(
			&node.NodeID,
			&node.ContentID,
			&node.UserID,
			&node.NotebookID,
			&node.DeletedAt,
			&node.CreatedAt,
			&node.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		nodeList = append(nodeList, &node)
	}
	return nodeList, nil
}
