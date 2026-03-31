package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/YagoSchramm/intensivo-first_service/model"
	"github.com/YagoSchramm/intensivo-first_service/repository"
	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	"github.com/google/uuid"
)

func TestRepositoryCRUDIntegration(t *testing.T) {
	couchURL := os.Getenv("COUCHDB_URL")
	if couchURL == "" {
		t.Skip("COUCHDB_URL not set; skipping integration test")
	}
	dbName := os.Getenv("COUCHDB_DB")
	if dbName == "" {
		u, _ := uuid.NewUUID()
		dbName = "test_notebook_" + u.String()
	}

	ctx := context.TODO()
	client, err := kivik.New("couch", couchURL)
	if err != nil {
		t.Fatalf("failed to create CouchDB client: %v", err)
	}
	created := true
	if err := client.CreateDB(ctx, dbName); err != nil {
		t.Fatalf("failed to create db: %v", err)
		created = false
	}
	defer func() {
		if created {
			_ = client.DestroyDB(ctx, dbName)
		}
	}()

	db := client.DB(dbName)
	repo := repository.NewRepository(db)

	docID := "nb-" + uuid.New().String()
	_, err = repo.Create(ctx, repository.NotebookDB{
		ID:          docID,
		Name:        "Notebook A",
		Description: "Desc",
		Rev:         "",
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	got, err := repo.Get(ctx, docID)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if got.Name != "Notebook A" || got.Description != "Desc" {
		t.Fatalf("unexpected doc: %#v", got)
	}

	_, err = repo.Update(ctx, model.Notebook{
		ID:          docID,
		Name:        "Notebook B",
		Description: "Desc 2",
	})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	_, err = repo.Delete(ctx, docID)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}
