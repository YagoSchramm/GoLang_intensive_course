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

func buildMetaContent(t *testing.T) (*service.MetaContentService, *service.NotebookService, func(), uuid.UUID, uuid.UUID) {
	t.Helper()

	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, _ := foundation.NewPostgresDB(conn)

	// Create user using service
	userRepo := repository.NewUserRepository(db)
	userSrv := service.NewUserService(userRepo)
	ctx := context.TODO()
	userInput := model.SignUpUserDTO{
		Name:     "Test Meta User",
		Email:    "meta@test.com",
		Password: "password123",
	}
	user, _ := userSrv.Create(ctx, userInput)
	var userId uuid.UUID
	if user != nil {
		userIdVal, _ := uuid.Parse(user.ID)
		userId = userIdVal
	}

	// Create notebook using service
	notebookRepo := repository.NewNotebookRepository(db)
	notebookSrv := service.NewNotebookService(notebookRepo)
	notebookInput := model.CreateNotebookDTO{
		UserID:      userId,
		Icon:        "folder",
		Name:        "Test Notebook",
		Image:       "google.com/image",
		Description: "description",
	}
	notebook, _ := notebookSrv.Create(ctx, notebookInput)
	var notebookId uuid.UUID
	if notebook != nil {
		notebookId = notebook.NotebookID
	}

	repo := repository.NewMetaContentRepository(db)
	srv := service.NewMetaContentService(repo)
	clean := func() {
		db.Close()
	}
	return srv, notebookSrv, clean, userId, notebookId
}

func TestServiceMetaContent(t *testing.T) {
	srv, _, clean, userId, notebookId := buildMetaContent(t)
	defer clean()

	t.Run("Create MetaContent", func(t *testing.T) {
		ctx := context.TODO()
		input := model.CreateMetaContentDTO{
			UserID:     userId,
			NotebookID: notebookId,
			Name:       "Test MetaContent",
			Icon:       "star",
		}
		metaContent, err := srv.Create(ctx, input)
		if err != nil {
			t.Fatalf("Erro na criação do MetaContent: %v", err)
		}
		if metaContent == nil {
			t.Fatalf("MetaContent não foi criado")
		}
		t.Log(metaContent)
	})

	t.Run("List MetaContent from User", func(t *testing.T) {
		ctx := context.TODO()
		input := model.ListMetaContentFromUserDTO{
			UserID: userId,
		}
		metaContents, err := srv.ListMetaContentFromUser(ctx, input)
		if err != nil {
			t.Fatalf("Erro ao listar MetaContent: %v", err)
		}
		if metaContents == nil || len(*metaContents) == 0 {
			t.Logf("Nenhum MetaContent encontrado para o usuário")
		} else {
			t.Log(metaContents)
		}
	})
}
