package service

import "github.com/YagoSchramm/intensivo-surfbook_v1/repository"

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
