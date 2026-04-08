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

type MetaContentHandler struct {
	srv *service.MetaContentService
}

func NewMetaContentHandler(s *service.MetaContentService) *MetaContentHandler {
	return &MetaContentHandler{srv: s}
}
func (h *MetaContentHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/metacontents", h.create).Methods("POST")
	r.HandleFunc("/metacontents", h.listMetaContents).Methods("GET")
	r.HandleFunc("/metacontents/{metacontent_id}", h.deleteMetaContent).Methods("DELETE")
	r.HandleFunc("/metacontents/{metacontent_id}", h.updateMetaContent).Methods("PATCH")
	r.HandleFunc("/metacontents/{metacontent_id}", h.findMetaContentById).Methods("GET")

}
func (h *MetaContentHandler) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	w.Header().Set("Content-Type", "application/json")
	var input model.CreateMetaContentDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input.UserID = userID
	resp, err := h.srv.Create(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao criar metacontent", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}
func (h *MetaContentHandler) listMetaContents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.ListMetaContentFromUserDTO{
		UserID: userID,
	}
	nbList, err := h.srv.ListMetaContentFromUser(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao listar metacontents: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(nbList); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}
func (h *MetaContentHandler) findMetaContentById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	metacontentID, err := uuid.Parse(vars["metacontent_id"])
	if err != nil {
		http.Error(w, "Header metacontent_id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.ListMetaContentByIdDTO{
		UserID:        userID,
		MetaContentID: metacontentID,
	}
	resp, err := h.srv.ListMetaContentById(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao buscar metacontent", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}
func (h *MetaContentHandler) deleteMetaContent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	metacontentID, err := uuid.Parse(vars["metacontent_id"])
	if err != nil {
		http.Error(w, "Header metacontent_id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.DeleteMetaContentDTO{
		MetaContentID: metacontentID,
		UserID:        userID,
	}
	err = h.srv.SoftDelete(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao deletar metacontent", http.StatusInternalServerError)
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
func (h *MetaContentHandler) updateMetaContent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	metacontentID, err := uuid.Parse(vars["metacontent_id"])
	if err != nil {
		http.Error(w, "Header metacontent_id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	var payload model.UpdateMetaContentDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	input := model.UpdateMetaContentDTO{
		MetaContentID: metacontentID,
		UserID:        userID,
		Name:          payload.Name,
		Icon:          payload.Icon,
	}
	resp, err := h.srv.Update(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao atualizar metacontent", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}
