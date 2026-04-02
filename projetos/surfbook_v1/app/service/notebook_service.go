package service

import (
	"github.com/YagoSchramm/intensivo-surfbook_v1/repository"
)

type NotebookService struct {
	repo *repository.NotebookRepository
}

func NewNotebookRepository(rep *repository.NotebookRepository) *NotebookService {
	return &NotebookService{repo: rep}
}
