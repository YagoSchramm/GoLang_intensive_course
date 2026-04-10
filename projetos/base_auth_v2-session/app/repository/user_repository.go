package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/YagoSchramm/base-auth-v2-session/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(d *sql.DB) *UserRepository {
	return &UserRepository{db: d}
}
func (r *UserRepository) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
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
