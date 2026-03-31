package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YagoSchramm/intensivo-first_service/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{service: srv}
}
func (h *Handler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/health", h.Health).Methods("GET")
	r.HandleFunc("/notebooks", h.Create).Methods("POST")
	r.HandleFunc("/notebooks/{notebook_id}", h.Get).Methods("GET")
	r.HandleFunc("/notebooks", h.Update).Methods("PUT")
	r.HandleFunc("/notebooks/{notebook_id}", h.Delete).Methods("DELETE")
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok", Message: "Everthing is cool here..."}); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
	}
}
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body service.CreateNotebookInput
	ctx := context.TODO()
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "falha ao decodificar corpo da requisicao", http.StatusInternalServerError)
	}
	resp, err := h.service.Create(ctx, body)
	if err != nil {
		http.Error(w, "falha ao inserir notebook no db", http.StatusInternalServerError)
	}
	println("DB response:", resp)
	w.Write([]byte("Created new notebook!"))
}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["notebook_id"]
	var nb service.Notebook
	ctx := context.TODO()
	nb, err := h.service.Get(ctx, id)
	if err != nil {
		http.Error(w, "falha ao consultar notebook", http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(nb); err != nil {
		http.Error(w, "Falha na decodificacao do notebook", http.StatusInternalServerError)
	}
}
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	var body service.Notebook
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Corpo invalido", http.StatusBadRequest)
		return
	}
	if body.ID == "" {
		http.Error(w, "O ID do notebook e obrigatorio", http.StatusBadRequest)
		return
	}
	newRev, err := h.service.Update(ctx, body)
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
		http.Error(w, "O ID do notebook e obrigatorio", http.StatusBadRequest)
		return
	}
	newRev, err := h.service.Delete(ctx, id)
	if err != nil {
		http.Error(w, "Falha ao deletar o notebook: "+err.Error(), http.StatusInternalServerError)
		return
	}
	println(newRev)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Deletado com sucesso"))
}
