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

type NodeContentHandler struct {
	srv *service.NodeContentService
}

func NewNodeContentHandler(srv *service.NodeContentService) *NodeContentHandler {
	return &NodeContentHandler{srv: srv}
}

func (h *NodeContentHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/node-contents", h.create).Methods("POST")
	r.HandleFunc("/node-contents", h.listNodeContents).Methods("GET")
	r.HandleFunc("/node-contents/{node_id}", h.deleteNodeContent).Methods("DELETE")
	r.HandleFunc("/node-contents/{node_id}", h.updateNodeContent).Methods("PATCH")
	r.HandleFunc("/node-contents/{node_id}", h.findNodeContentById).Methods("GET")
}

func (h *NodeContentHandler) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	w.Header().Set("Content-Type", "application/json")

	var input model.CreateNodeContentDTO
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
		http.Error(w, "Erro ao criar node content", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (h *NodeContentHandler) listNodeContents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.ListNodeContentFromUserDTO{
		UserID: userID,
	}
	resp, err := h.srv.ListFromUser(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao listar node contents: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (h *NodeContentHandler) findNodeContentById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	nodeID, err := uuid.Parse(vars["node_id"])
	if err != nil {
		http.Error(w, "Header node_id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.FindNodeContentFromUserDTO{
		UserID: userID,
		NodeID: nodeID,
	}
	resp, err := h.srv.FindByUserNodeId(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao buscar node content", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func (h *NodeContentHandler) deleteNodeContent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	nodeID, err := uuid.Parse(vars["node_id"])
	if err != nil {
		http.Error(w, "Header node_id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := model.DeleteNodeContentDTO{
		UserID: userID,
		NodeID: nodeID,
	}
	err = h.srv.SoftDelete(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao deletar node content", http.StatusInternalServerError)
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

func (h *NodeContentHandler) updateNodeContent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	nodeID, err := uuid.Parse(vars["node_id"])
	if err != nil {
		http.Error(w, "Header node_id invÃ¡lido: "+err.Error(), http.StatusBadRequest)
		return
	}
	var payload model.UpdateNodeContentDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invÃ¡lido", http.StatusBadRequest)
		return
	}
	input := model.UpdateNodeContentDTO{
		UserID:     userID,
		NodeID:     nodeID,
		ContentID:  payload.ContentID,
		NotebookID: payload.NotebookID,
	}
	resp, err := h.srv.Update(ctx, input)
	if err != nil {
		http.Error(w, "Erro ao atualizar node content", http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}
