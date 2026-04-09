package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/YagoSchramm/base-auth-v1/foundation"
	"github.com/YagoSchramm/base-auth-v1/model"
	"github.com/YagoSchramm/base-auth-v1/repository"
	"github.com/go-kivik/kivik/v4"
)

type AuthenticationService struct {
	repo *repository.Repository
}

func NewAuthenticationService(r *repository.Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) SignUp(ctx context.Context, input model.SignUpUserDTO) (*model.UserAcces, error) {

	new_user := &model.UserEntityDomain{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	userNameExist, err := s.repo.GetUser(ctx, new_user.Name)
	if err != nil && kivik.HTTPStatus(err) != http.StatusNotFound {
		return nil, errors.New("Internal error")
	}
	if userNameExist != nil {
		return nil, errors.New("Invalid Credentials")
	}
	_, err = s.repo.CreateUser(ctx, &input)
	if err != nil {
		return nil, err
	}
	foundation.SendMockEmail(input.Email, "Welcome!", "SingUp on our Service")
	return &model.UserAcces{
		User:  *new_user,
		Token: foundation.ToBase64(new_user.Name),
	}, nil

}
func (s *Service) SignIn(ctx context.Context, input model.SignInUserDTO) (*model.UserAcces, error) {

	current_user, err := s.repo.GetUser(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if current_user.Password != input.Password {
		return nil, errors.New(" Invalid Credential")
	}
	foundation.SendMockEmail(input.Name, "New SignIn", "New SignIn in your accont")

	return &model.UserAcces{
		User:  *current_user,
		Token: foundation.ToBase64(current_user.Name),
	}, nil
}

func (s *Service) Finduser(ctx context.Context, username string) (*model.UserEntityDomain, error) {

	current_user, err := s.repo.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	return current_user, nil
}
