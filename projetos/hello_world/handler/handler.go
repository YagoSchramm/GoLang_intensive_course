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
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	var body Notebook
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Corpo inválido", http.StatusBadRequest)
		return
	}

	var existing Notebook
	err := h.couchdb.Get(ctx, body.ID).ScanDoc(&existing)
	if err != nil {
		http.Error(w, "Notebook não encontrado", http.StatusNotFound)
		return
	}
	body.Rev = existing.Rev
	newRev, err := h.couchdb.Put(ctx, body.ID, body)
	if err != nil {
		http.Error(w, "Erro ao atualizar Notebook: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	println(newRev)
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	vars := mux.Vars(r)
	id := vars["notebook_id"]
	if id == "" {
		http.Error(w, "O ID do notebook é obrigatório", http.StatusBadRequest)
		return
	}

	// 2. Buscar a revisão atual (_rev)
	// No CouchDB, para deletar, você precisa do ID + REV.
	var nb Notebook
	err := h.couchdb.Get(ctx, id).ScanDoc(&nb)
	if err != nil {
		http.Error(w, "Notebook não encontrado", http.StatusNotFound)
		return
	}
	newRev, err := h.couchdb.Delete(ctx, id, nb.Rev)
	if err != nil {
		http.Error(w, "Falha ao deletar o notebook: "+err.Error(), http.StatusInternalServerError)
		return
	}

	println(newRev)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Deletado com sucesso"))
}
