package service

import (
	"context"

	"github.com/go-kivik/kivik/v4"
	"github.com/google/uuid"
)

type Service struct {
	db *kivik.DB
}
type CreateNotebookInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Notebook struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type NotebookDB struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rev         string `json:"_rev,omitempty"`
}

func NewService(d *kivik.DB) *Service {
	return &Service{db: d}
}

func (s *Service) Create(ctx context.Context, input CreateNotebookInput) (string, error) {
	uid, _ := uuid.NewUUID()
	id := uid.String()
	nb := NotebookDB{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Rev:         "",
	}
	resp, err := s.db.Put(ctx, id, nb)
	if err != nil {
		return "", err
	}
	return resp, nil
}
func (s *Service) Get(ctx context.Context, id string) (Notebook, error) {
	var nb NotebookDB
	err := s.db.Get(ctx, id).ScanDoc(&nb)
	if err != nil {
		return Notebook{}, err
	}
	new_nb := Notebook{
		ID:          nb.ID,
		Name:        nb.Name,
		Description: nb.Description,
	}
	return new_nb, err
}

func (s *Service) Update(ctx context.Context, input Notebook) (string, error) {
	var existing NotebookDB
	if err := s.db.Get(ctx, input.ID).ScanDoc(&existing); err != nil {
		return "", err
	}
	updated := NotebookDB{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Rev:         existing.Rev,
	}
	resp, err := s.db.Put(ctx, input.ID, updated)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (s *Service) Delete(ctx context.Context, id string) (string, error) {
	var existing NotebookDB
	if err := s.db.Get(ctx, id).ScanDoc(&existing); err != nil {
		return "", err
	}
	resp, err := s.db.Delete(ctx, id, existing.Rev)
	if err != nil {
		return "", err
	}
	return resp, nil
}
