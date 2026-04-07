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
	r.HandleFunc("/notebooks/{notebook_id}", h.deleteNotebook).Methods("DELETE")
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
func (h *NotebookHandler) deleteNotebook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	notebookID, err := uuid.Parse(vars["user-id"])
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.DeleteNotebookDTO{
		NotebookID: notebookID,
		UserID:     userID,
	}
	err = h.srv.SoftDelete(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao deletar notebook", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Operação realizada com sucesso",
	})
	if err != nil {
		http.Error(w, "Operação realizada e Erro interno: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
