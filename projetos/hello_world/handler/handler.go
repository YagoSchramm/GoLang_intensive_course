package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kivik/kivik/v4"
	"github.com/gorilla/mux"
)

type Handler struct {
	couchdb *kivik.DB
}

func New(c *kivik.DB) *Handler {
	return &Handler{couchdb: c}
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Notebook struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rev         string `json:"_rev,omitempty"`
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok", Message: "Everthing is cool here..."}); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
	}
}
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body Notebook
	ctx := context.TODO()
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "falha ao decodificar corpo da requisicao", http.StatusInternalServerError)
	}
	resp, err := h.couchdb.Put(ctx, body.ID, body)
	if err != nil {
		http.Error(w, "falha ao criar notebook", http.StatusInternalServerError)
	}
	println("DB response:", resp)
	w.Write([]byte(body.ID))
}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["notebook_id"]
	var nb Notebook
	ctx := context.TODO()
	err := h.couchdb.Get(ctx, id).ScanDoc(&nb)
	if err != nil {
		http.Error(w, "falha ao consultar notebook", http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(nb); err != nil {
		http.Error(w, "Falha na decodificação do notebook", http.StatusInternalServerError)
	}
}
