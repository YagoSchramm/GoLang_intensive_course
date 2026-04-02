package service_test

import (
	"testing"

	"github.com/YagoSchramm/intensivo-surfbook_v1/foundation"
	"github.com/YagoSchramm/intensivo-surfbook_v1/repository"
	"github.com/YagoSchramm/intensivo-surfbook_v1/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func build(t *testing.T) (*service.UserService, func()) {
	t.Helper()
	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, _ := foundation.NewPostgresDB(conn)
	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	return srv, t.Cleanup()
}
