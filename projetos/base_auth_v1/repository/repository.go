package repository

import (
	"context"
	"log"

	"github.com/YagoSchramm/base-auth-v1/model"
	"github.com/go-kivik/kivik/v4"
)

type Repository struct {
	db     *kivik.DB
	userdb *kivik.DB
}

type NotebookDB struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rev         string `json:"_rev,omitempty"`
}
type UserDB struct {
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Rev      string `json:"_rev,omitempty"`
}

func NewRepository(d *kivik.DB, u *kivik.DB) *Repository {
	return &Repository{db: d, userdb: u}
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
func (r *Repository) CreateUser(ctx context.Context, u *model.SignUpUserDTO) (string, error) {
	user := UserDB{
		UserName: u.Name,
		Email:    u.Email,
		Password: u.Password,
		Rev:      "",
	}
	_, err := r.userdb.Put(ctx, u.Name, user)
	if err != nil {
		return "", err
	}
	return "", err
}

func (r *Repository) GetUser(ctx context.Context, userName string) (*model.UserEntityDomain, error) {
	var user UserDB
	if err := r.userdb.Get(ctx, userName).ScanDoc(&user); err != nil {
		log.Print(err)
		return nil, err
	}
	userdto := model.UserEntityDomain{
		Name:     user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	return &userdto, nil
}
