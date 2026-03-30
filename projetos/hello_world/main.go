package main

import (
	"log"
	"net/http"

	"github.com/YagoSchramm/intensivo-helloWorld/projetos/hello_world/handler"
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

	log.Println("Basic web server")
	r := mux.NewRouter()
	h := handler.New(db)
	r.HandleFunc("/health", h.Health).Methods("GET")
	r.HandleFunc("/notebooks", h.Create).Methods("POST")
	r.HandleFunc("/notebooks/{notebook_id}", h.Get).Methods("GET")
	http.ListenAndServe(":8000", r)
}
