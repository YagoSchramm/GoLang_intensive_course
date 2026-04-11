package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YagoSchramm/base-auth-v2-session/foundation"
	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/YagoSchramm/base-auth-v2-session/repository"
	"github.com/YagoSchramm/base-auth-v2-session/service"
	"github.com/google/uuid"
)

func buildUser(t *testing.T) (*service.UserService, func()) {
	t.Helper()

	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, _ := foundation.NewPostgresDB(conn)
	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	clean := func() {
		db.Close()
	}
	return srv, clean
}

func TestServiceUser(t *testing.T) {
	srv, clean := buildUser(t)
	defer clean()

	t.Run("Create User with SignUp", func(t *testing.T) {
		ctx := context.TODO()
		uniqueEmail := fmt.Sprintf("newuser%s@test.com", uuid.New().String()[:8])
		input := model.SignUpUserDTO{
			Name:     "New User",
			Email:    uniqueEmail,
			Password: "securepassword123",
		}
		user, err := srv.Create(ctx, input)
		if err != nil {
			t.Fatalf("Erro na criação do usuário: %v", err)
		}
		if user == nil {
			t.Fatalf("Usuário não foi criado")
		}
		if user.Email != input.Email {
			t.Fatalf("Email não corresponde: esperado %s, obtido %s", input.Email, user.Email)
		}
		t.Log(user)
	})

	t.Run("Authenticate User with SignIn", func(t *testing.T) {
		ctx := context.TODO()

		uniqueEmail := fmt.Sprintf("auth%s@test.com", uuid.New().String()[:8])
		signUpInput := model.SignUpUserDTO{
			Name:     "Auth User",
			Email:    uniqueEmail,
			Password: "password123",
		}
		user, err := srv.Create(ctx, signUpInput)
		if err != nil {
			t.Fatalf("Erro na criação do usuário: %v", err)
		}

		// Try to authenticate
		signInInput := model.SignInUserDTO{
			Email:    uniqueEmail,
			Password: "password123",
		}
		authenticatedUser, err := srv.Authenticate(ctx, signInInput)
		if err != nil {
			t.Fatalf("Erro na autenticação do usuário: %v", err)
		}
		if authenticatedUser == nil {
			t.Fatalf("Usuário não foi autenticado")
		}
		if authenticatedUser.Email != user.Email {
			t.Fatalf("Email não corresponde: esperado %s, obtido %s", user.Email, authenticatedUser.Email)
		}
		t.Log(authenticatedUser)
	})

	t.Run("Authenticate User with wrong password", func(t *testing.T) {
		ctx := context.TODO()

		uniqueEmail := fmt.Sprintf("wrongpass%s@test.com", uuid.New().String()[:8])
		signUpInput := model.SignUpUserDTO{
			Name:     "Wrong Pass User",
			Email:    uniqueEmail,
			Password: "password123",
		}
		_, err := srv.Create(ctx, signUpInput)
		if err != nil {
			t.Fatalf("Erro na criação do usuário: %v", err)
		}

		// Try to authenticate with wrong password
		signInInput := model.SignInUserDTO{
			Email:    uniqueEmail,
			Password: "wrongpassword",
		}
		authenticatedUser, err := srv.Authenticate(ctx, signInInput)
		if err != nil {
			t.Fatalf("Erro na autenticação do usuário: %v", err)
		}
		if authenticatedUser != nil {
			t.Fatalf("Usuário não deveria ter sido autenticado com senha incorreta")
		}
		t.Log("Autenticação falhou como esperado")
	})
}
