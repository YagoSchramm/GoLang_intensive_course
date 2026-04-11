package service

import (
	"context"

	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/YagoSchramm/base-auth-v2-session/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (srv *UserService) Create(ctx context.Context, input model.SignUpUserDTO) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:       uuid.NewString(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	return srv.repo.CreateUser(ctx, &user)
}

func (srv *UserService) Authenticate(ctx context.Context, input model.SignInUserDTO) (*model.User, error) {
	user, err := srv.repo.GetUser(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, nil
	}

	return user, nil
}

func (srv *UserService) FindByID(ctx context.Context, id string) (*model.User, error) {
	return srv.repo.FindByID(ctx, id)
}

func (srv *UserService) List(ctx context.Context) ([]*model.User, error) {
	return srv.repo.ListUsers(ctx)
}

func (srv *UserService) Update(ctx context.Context, input model.User) (*model.User, error) {
	if err := srv.repo.UpdateUser(ctx, &input); err != nil {
		return nil, err
	}
	return srv.repo.FindByID(ctx, input.ID)
}

func (srv *UserService) Delete(ctx context.Context, id string) error {
	return srv.repo.DeleteUser(ctx, id)
}
