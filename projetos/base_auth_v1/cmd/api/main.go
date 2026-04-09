package main

import (
	"log"
	"net/http"

	"github.com/YagoSchramm/base-auth-v1/handler"
	"github.com/YagoSchramm/base-auth-v1/repository"
	"github.com/YagoSchramm/base-auth-v1/service"
	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	"github.com/gorilla/mux"
)

func main() {
	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("erro ao criar cliente CouchDB: %s", err)
	}
	db := client.DB("notebook")
	userdb := client.DB("user")
	log.Println("Basic web server")
	repo := repository.NewRepository(db, userdb)
	srv := service.NewService(repo)
	r := mux.NewRouter()
	h := handler.NewHandler(srv)
	h.MountHandlers(r)
	http.ListenAndServe(":8000", r)
}
