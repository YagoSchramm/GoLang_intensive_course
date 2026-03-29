package handler

import "net/http"

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
func HandlerPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}
