package service

import (
	"context"
	"time"

	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/YagoSchramm/base-auth-v2-session/repository"
	"github.com/google/uuid"
)

type NodeContentService struct {
	repo *repository.NodeContentRepository
}

func NewNodeContentService(r *repository.NodeContentRepository) *NodeContentService {
	return &NodeContentService{repo: r}
}

func (srv *NodeContentService) Create(ctx context.Context, input model.CreateNodeContentDTO) (*model.NodeContentEntity, error) {
	id := uuid.New()
	now := time.Now()
	node := model.NodeContentEntity{
		NodeID:     id,
		ContentID:  input.ContentID,
		UserID:     input.UserID,
		NotebookID: input.NotebookID,
		DeletedAt:  nil,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	err := srv.repo.Create(ctx, &node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (srv *NodeContentService) ListFromUser(ctx context.Context, input model.ListNodeContentFromUserDTO) ([]*model.NodeContentEntity, error) {
	return srv.repo.ListNodeContents(ctx, input)
}

func (srv *NodeContentService) FindByUserNodeId(ctx context.Context, input model.FindNodeContentFromUserDTO) (*model.NodeContentEntity, error) {
	return srv.repo.FindByUserIdNodeId(ctx, input)
}

func (srv *NodeContentService) SoftDelete(ctx context.Context, input model.DeleteNodeContentDTO) error {
	return srv.repo.Delete(ctx, input)
}

func (srv *NodeContentService) Update(ctx context.Context, input model.UpdateNodeContentDTO) (*model.NodeContentEntity, error) {
	err := srv.repo.Update(ctx, input)
	if err != nil {
		return nil, err
	}
	find := model.FindNodeContentFromUserDTO{
		UserID: input.UserID,
		NodeID: input.NodeID,
	}
	node, err := srv.repo.FindByUserIdNodeId(ctx, find)
	if err != nil {
		return nil, err
	}
	return node, nil
}
