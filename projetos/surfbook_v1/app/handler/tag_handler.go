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

type TagHandler struct {
	srv *service.TagService
}

func NewTagHandler(srv *service.TagService) *TagHandler {
	return &TagHandler{srv: srv}
}

func (h *TagHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/tags", h.create).Methods("POST")
	r.HandleFunc("/tags", h.listTags).Methods("GET")
	r.HandleFunc("/tags/{tag_id}", h.deleteTag).Methods("DELETE")
	r.HandleFunc("/tags/{tag_id}", h.updateTag).Methods("PATCH")
	r.HandleFunc("/tags/{tag_id}", h.findTagById).Methods("GET")
}

func (h *TagHandler) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	w.Header().Set("Content-Type", "application/json")

	var input model.CreateTagDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON invÃ¡lido", http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input.UserID = userID

	resp, err := h.srv.Create(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao criar tag", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (h *TagHandler) listTags(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.ListTagsFromUserDTO{
		UserID: userID,
	}
	resp, err := h.srv.ListFromUser(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao listar tags: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (h *TagHandler) findTagById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	tagID, err := uuid.Parse(vars["tag_id"])
	if err != nil {
		http.Error(w, "Header tag_id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.FindTagFromUserDTO{
		UserID: userID,
		TagID:  tagID,
	}
	resp, err := h.srv.FindByUserTagId(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao buscar tag", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (h *TagHandler) deleteTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	tagID, err := uuid.Parse(vars["tag_id"])
	if err != nil {
		http.Error(w, "Header tag_id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.DeleteTagDTO{
		UserID: userID,
		TagID:  tagID,
	}
	err = h.srv.SoftDelete(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao deletar tag", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "OperaÃ§Ã£o realizada com sucesso",
	})
	if err != nil {
		http.Error(w, "OperaÃ§Ã£o realizada e Erro interno: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TagHandler) updateTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	tagID, err := uuid.Parse(vars["tag_id"])
	if err != nil {
		http.Error(w, "Header tag_id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	var payload model.UpdateTagDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invÃ¡lido", http.StatusBadRequest)
		return
	}
	input := model.UpdateTagDTO{
		UserID: userID,
		TagID:  tagID,
		Name:   payload.Name,
		Color:  payload.Color,
	}
	resp, err := h.srv.Update(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao atualizar tag", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}
