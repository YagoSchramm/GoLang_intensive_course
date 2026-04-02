package service

import (
	"context"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
	"github.com/YagoSchramm/intensivo-surfbook_v1/repository"
)

type NotebookService struct {
	repo *repository.NotebookRepository
}

func NewNotebookService(rep *repository.NotebookRepository) *NotebookService {
	return &NotebookService{repo: rep}
}
func (n *NotebookService) Create(ctx context.Context, input model.CreateNotebookDTO) (*model.NotebookEntity, error) {

	err := n.repo.Create(ctx, notebook)
}
