package service_test

import (
	"context"
	"testing"

	"github.com/YagoSchramm/base-auth-v2-session/foundation"
	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/YagoSchramm/base-auth-v2-session/repository"
	"github.com/YagoSchramm/base-auth-v2-session/service"
	"github.com/google/uuid"
)

func build(t *testing.T) (*service.NotebookService, func(), uuid.UUID) {
	t.Helper()

	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, _ := foundation.NewPostgresDB(conn)

	// Create user using service
	userRepo := repository.NewUserRepository(db)
	userSrv := service.NewUserService(userRepo)
	ctx := context.TODO()
	userInput := model.SignUpUserDTO{
		Name:     "Test Notebook User",
		Email:    "notebook@test.com",
		Password: "password123",
	}
	user, _ := userSrv.Create(ctx, userInput)
	var userId uuid.UUID
	if user != nil {
		userIdVal, _ := uuid.Parse(user.ID)
		userId = userIdVal
	}

	repo := repository.NewNotebookRepository(db)
	srv := service.NewNotebookService(repo)
	clean := func() {
		db.Close()
	}
	return srv, clean, userId
}
func TestServiceNotebook(t *testing.T) {
	srv, clean, userId := build(t)
	defer clean()
	t.Run("Create Notebook", func(t *testing.T) {
		ctx := context.TODO()
		input := model.CreateNotebookDTO{
			UserID:      userId,
			Icon:        "folder",
			Name:        "Test Notebook",
			Image:       "google.com/image",
			Description: "description notebook",
		}
		notebook, err := srv.Create(ctx, input)
		if err != nil {
			t.Fatalf("Erro na criação do Notebook:%v", err)
		}
		t.Log(notebook)
	})

}
