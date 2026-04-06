package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YagoSchramm/intensivo-surfbook_v1/foundation"
	"github.com/YagoSchramm/intensivo-surfbook_v1/handler"
	"github.com/YagoSchramm/intensivo-surfbook_v1/repository"
	"github.com/YagoSchramm/intensivo-surfbook_v1/service"
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
	notebookRepo := repository.NewNotebookRepository(db)
	notebookSrv := service.NewNotebookService(notebookRepo)
	notebookHandler := handler.NewNotebookHandler(notebookSrv)
	r := mux.NewRouter()
	notebookHandler.MountHandlers(r)
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Inicialização do servidor não executada:%v ", err)
	}
}
