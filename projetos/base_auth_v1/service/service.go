package service

import (
	"context"

	"github.com/YagoSchramm/intensivo-first_service/model"
	"github.com/YagoSchramm/intensivo-first_service/repository"
	"github.com/google/uuid"
)

type Service struct {
	repo NotebookRepository
}

type NotebookRepository interface {
	Create(ctx context.Context, nb repository.NotebookDB) (string, error)
	Get(ctx context.Context, id string) (repository.NotebookDB, error)
	Update(ctx context.Context, input model.Notebook) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}

func NewService(r NotebookRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, input model.CreateNotebookInput) (string, error) {
	uid, _ := uuid.NewUUID()
	id := uid.String()
	nb := repository.NotebookDB{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Rev:         "",
	}
	return s.repo.Create(ctx, nb)
}
func (s *Service) Get(ctx context.Context, id string) (model.Notebook, error) {
	nb, err := s.repo.Get(ctx, id)
	if err != nil {
		return model.Notebook{}, err
	}
	new_nb := model.Notebook{
		ID:          id,
		Name:        nb.Name,
		Description: nb.Description,
	}
	return new_nb, err
}

func (s *Service) Update(ctx context.Context, input model.Notebook) (string, error) {
	return s.repo.Update(ctx, input)
}

func (s *Service) Delete(ctx context.Context, id string) (string, error) {
	return s.repo.Delete(ctx, id)
}
