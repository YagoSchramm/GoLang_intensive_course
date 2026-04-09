package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/YagoSchramm/base-auth-v1/model"
	"github.com/YagoSchramm/base-auth-v1/repository"
	"github.com/YagoSchramm/base-auth-v1/service"
	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func buildService(t *testing.T) (*service.Service, *repository.Repository, func()) {
	t.Helper()

	couchURL := os.Getenv("COUCHDB_URL")
	if couchURL == "" {
		couchURL = "http://admin:pass@localhost:5984/"
	}

	dbName := os.Getenv("COUCHDB_DB")
	if dbName == "" {
		dbName = "test_notebook_" + uuid.New().String()
	}

	ctx := context.TODO()
	client, err := kivik.New("couch", couchURL)
	if err != nil {
		t.Skipf("CouchDB indisponivel: %v", err)
	}
	created := true
	if err := client.CreateDB(ctx, dbName); err != nil {
		t.Skipf("CouchDB indisponivel para criar DB: %v", err)
		created = false
	}

	db := client.DB(dbName)
	repo := repository.NewRepository(db, db)
	srv := service.NewService(repo)

	cleanup := func() {
		if created {
			_ = client.DestroyDB(ctx, dbName)
		}
	}
	return srv, repo, cleanup
}

func TestServiceCreate(t *testing.T) {
	srv, _, cleanup := buildService(t)
	defer cleanup()

	resp, err := srv.Create(context.TODO(), model.CreateNotebookInput{
		Name:        "Notebook A",
		Description: "Desc",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestServiceGet(t *testing.T) {
	srv, repo, cleanup := buildService(t)
	defer cleanup()

	id := "nb-" + uuid.New().String()
	_, err := repo.Create(context.TODO(), repository.NotebookDB{
		ID:          id,
		Name:        "Notebook A",
		Description: "Desc",
		Rev:         "",
	})
	assert.NoError(t, err)

	nb, err := srv.Get(context.TODO(), id)
	assert.NoError(t, err)
	assert.Equal(t, id, nb.ID)
	assert.Equal(t, "Notebook A", nb.Name)
	assert.Equal(t, "Desc", nb.Description)
}

func TestServiceUpdate(t *testing.T) {
	srv, repo, cleanup := buildService(t)
	defer cleanup()

	id := "nb-" + uuid.New().String()
	_, err := repo.Create(context.TODO(), repository.NotebookDB{
		ID:          id,
		Name:        "Notebook A",
		Description: "Desc",
		Rev:         "",
	})
	assert.NoError(t, err)

	resp, err := srv.Update(context.TODO(), model.Notebook{
		ID:          id,
		Name:        "Notebook B",
		Description: "New Desc",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestServiceDelete(t *testing.T) {
	srv, repo, cleanup := buildService(t)
	defer cleanup()

	id := "nb-" + uuid.New().String()
	_, err := repo.Create(context.TODO(), repository.NotebookDB{
		ID:          id,
		Name:        "Notebook A",
		Description: "Desc",
		Rev:         "",
	})
	assert.NoError(t, err)

	resp, err := srv.Delete(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
