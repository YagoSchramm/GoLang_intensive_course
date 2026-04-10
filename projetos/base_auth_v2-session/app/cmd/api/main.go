package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/YagoSchramm/base-auth-v2-session/foundation"
	"github.com/YagoSchramm/base-auth-v2-session/handler"
	"github.com/YagoSchramm/base-auth-v2-session/middleware"
	"github.com/YagoSchramm/base-auth-v2-session/repository"
	"github.com/YagoSchramm/base-auth-v2-session/service"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("creating a webserver")
	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, err := foundation.NewPostgresDB(conn)
	if err != nil {
		log.Fatalf("Conexão com PostgreSQL não executada!: %v", err)
	}
	fmt.Println("Conexão com PostgreSQL estabelecida com sucesso.")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret"
	}

	userRepo := repository.NewUserRepository(db)
	userSrv := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSrv, jwtSecret)

	notebookRepo := repository.NewNotebookRepository(db)
	notebookSrv := service.NewNotebookService(notebookRepo)
	notebookHandler := handler.NewNotebookHandler(notebookSrv, jwtSecret)
	tagRepo := repository.NewTagRepository(db)
	tagSrv := service.NewTagService(tagRepo)
	tagHandler := handler.NewTagHandler(tagSrv, jwtSecret)
	nodeContentRepo := repository.NewNodeContentRepository(db)
	nodeContentSrv := service.NewNodeContentService(nodeContentRepo)
	nodeContentHandler := handler.NewNodeContentHandler(nodeContentSrv, jwtSecret)

	r := mux.NewRouter()
	userHandler.MountHandlers(r)

	protected := r.NewRoute().Subrouter()
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	notebookHandler.MountHandlers(protected)
	tagHandler.MountHandlers(protected)
	nodeContentHandler.MountHandlers(protected)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Inicialização do servidor não executada:%v ", err)
	}
}
