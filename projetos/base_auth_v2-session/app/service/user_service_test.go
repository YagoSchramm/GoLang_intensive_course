package service_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/YagoSchramm/base-auth-v2-session/foundation"
	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/YagoSchramm/base-auth-v2-session/repository"
	"github.com/YagoSchramm/base-auth-v2-session/service"
	"github.com/google/uuid"
)

func createUser(db *sql.DB, id uuid.UUID, email, password string) error {
	query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, id, email, password)
	return err
}

func build(t *testing.T) (*service.NotebookService, func(), uuid.UUID) {
	t.Helper()
	userId := uuid.New()

	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, _ := foundation.NewPostgresDB(conn)
	createUser(db, userId, "test@test.com", "password123")
	repo := repository.NewNotebookRepository(db)
	srv := service.NewNotebookService(repo)
	clean := func() {
		db.Close()
	}
	return srv, clean, userId
}
func TestServiceNotebook(t *testing.T) {
	srv, _, userId := build(t)
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
