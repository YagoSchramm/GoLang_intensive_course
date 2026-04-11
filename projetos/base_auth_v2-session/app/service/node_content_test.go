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

func buildNodeContent(t *testing.T) (*service.NodeContentService, *service.MetaContentService, *service.NotebookService, func(), uuid.UUID, uuid.UUID, uuid.UUID) {
	t.Helper()

	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, _ := foundation.NewPostgresDB(conn)

	// Create user using service
	userRepo := repository.NewUserRepository(db)
	userSrv := service.NewUserService(userRepo)
	ctx := context.TODO()
	userInput := model.SignUpUserDTO{
		Name:     "Test Node User",
		Email:    "node@test.com",
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

	// Create meta_content using service
	metaRepo := repository.NewMetaContentRepository(db)
	metaSrv := service.NewMetaContentService(metaRepo)
	metaInput := model.CreateMetaContentDTO{
		UserID:     userId,
		NotebookID: notebookId,
		Name:       "Test Meta Content",
		Icon:       "star",
	}
	metaContent, _ := metaSrv.Create(ctx, metaInput)
	var metaContentId uuid.UUID
	if metaContent != nil {
		metaContentId = metaContent.MetaContentID
	}

	repo := repository.NewNodeContentRepository(db)
	srv := service.NewNodeContentService(repo)
	clean := func() {
		db.Close()
	}
	return srv, metaSrv, notebookSrv, clean, userId, notebookId, metaContentId
}

func TestServiceNodeContent(t *testing.T) {
	srv, _, _, clean, userId, notebookId, metaContentId := buildNodeContent(t)
	defer clean()

	t.Run("Create NodeContent", func(t *testing.T) {
		ctx := context.TODO()
		input := model.CreateNodeContentDTO{
			UserID:     userId,
			ContentID:  metaContentId,
			NotebookID: notebookId,
		}
		nodeContent, err := srv.Create(ctx, input)
		if err != nil {
			t.Fatalf("Erro na criação do NodeContent: %v", err)
		}
		if nodeContent == nil {
			t.Fatalf("NodeContent não foi criado")
		}
		t.Log(nodeContent)
	})

	t.Run("List NodeContent from User", func(t *testing.T) {
		ctx := context.TODO()
		input := model.ListNodeContentFromUserDTO{
			UserID: userId,
		}
		nodeContents, err := srv.ListFromUser(ctx, input)
		if err != nil {
			t.Fatalf("Erro ao listar NodeContent: %v", err)
		}
		if len(nodeContents) == 0 {
			t.Logf("Nenhum NodeContent encontrado para o usuário")
		} else {
			t.Log(nodeContents)
		}
	})
}
