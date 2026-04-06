package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YagoSchramm/intensivo-surfbook_v1/model"
	"github.com/YagoSchramm/intensivo-surfbook_v1/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type NotebookHandler struct {
	srv *service.NotebookService
}

func NewNotebookHandler(srv *service.NotebookService) *NotebookHandler {
	return &NotebookHandler{srv: srv}
}
func (h *NotebookHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/notebooks", h.create).Methods("POST")
	r.HandleFunc("/notebooks", h.listNotebooks).Methods("GET")
}
func (h *NotebookHandler) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	w.Header().Set("Content-Type", "application/json")
	var input model.CreateNotebookDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	resp, err := h.srv.Create(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao criar notebook", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}
func (h *NotebookHandler) listNotebooks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.ListNotebookFromUserDTO{
		User_id: userID,
	}
	nbList, err := h.srv.ListFromUser(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao listar notebooks: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(nbList); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}
