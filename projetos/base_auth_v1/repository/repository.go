package repository

import (
	"context"

	"github.com/YagoSchramm/intensivo-first_service/model"
	"github.com/go-kivik/kivik/v4"
)

type Repository struct {
	db *kivik.DB
}

type NotebookDB struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rev         string `json:"_rev,omitempty"`
}

func NewRepository(d *kivik.DB) *Repository {
	return &Repository{db: d}
}

func (r *Repository) Create(ctx context.Context, nb NotebookDB) (string, error) {
	return r.db.Put(ctx, nb.ID, nb)
}

func (r *Repository) Get(ctx context.Context, id string) (NotebookDB, error) {
	var nb NotebookDB
	if err := r.db.Get(ctx, id).ScanDoc(&nb); err != nil {
		return NotebookDB{}, err
	}
	return nb, nil
}

func (r *Repository) Update(ctx context.Context, input model.Notebook) (string, error) {
	existing, err := r.Get(ctx, input.ID)
	if err != nil {
		return "", err
	}
	updated := NotebookDB{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Rev:         existing.Rev,
	}
	return r.db.Put(ctx, input.ID, updated)
}

func (r *Repository) Delete(ctx context.Context, id string) (string, error) {
	existing, err := r.Get(ctx, id)
	if err != nil {
		return "", err
	}
	return r.db.Delete(ctx, id, existing.Rev)
}
