package main

import (
	"log"
	"net/http"

	"github.com/YagoSchramm/intensivo-helloWorld/projetos/hello_world/handler"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Basic web server")
	r := mux.NewRouter()
	r.HandleFunc("/helloworld", handler.HandlerHello).Methods("GET")
	r.HandleFunc("/ping", handler.HandlerPing).Methods("POST")
	http.ListenAndServe(":8000", r)
}
