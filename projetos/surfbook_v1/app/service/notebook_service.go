package service

import (
	"context"
	"time"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
	"github.com/YagoSchramm/intensivo-surfbook_v1/repository"
	"github.com/google/uuid"
)

type NotebookService struct {
	repo *repository.NotebookRepository
}

func NewNotebookService(rep *repository.NotebookRepository) *NotebookService {
	return &NotebookService{repo: rep}
}
func (n *NotebookService) Create(ctx context.Context, input model.CreateNotebookDTO) (*model.NotebookEntity, error) {
	id := uuid.New()
	now := time.Now()
	notebook := model.NotebookEntity{
		NotebookID:  id,
		UserID:      input.UserID,
		Name:        input.Name,
		Image:       input.Image,
		Description: input.Description,
		Icon:        input.Icon,
		DeletedAt:   nil,
		UpdatedAt:   now,
		CreatedAt:   now,
	}
	err := n.repo.Create(ctx, &notebook)
	if err != nil {
		return nil, err
	}

	return &notebook, nil
}
func (srv *NotebookService) ListFromUser(ctx context.Context, input model.ListNotebookFromUserDTO) ([]*model.NotebookEntity, error) {
	return srv.repo.ListNotebooks(ctx, input.User_id)
}
func (srv *NotebookService) SoftDelete(ctx context.Context, input model.DeleteNotebookDTO) error {
	return srv.repo.Delete(ctx, input)
}
