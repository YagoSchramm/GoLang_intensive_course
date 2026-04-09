package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YagoSchramm/base-auth-v1/middleware"
	"github.com/YagoSchramm/base-auth-v1/model"
	"github.com/YagoSchramm/base-auth-v1/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{service: srv}
}
func (h *Handler) MountHandlers(r *mux.Router) {
	identifyMiddleware := middleware.Identify(h.service.Finduser)
	r.Use(middleware.LogStartandDuration)
	r.HandleFunc("/health", h.Health).Methods("GET")
	a := r.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/signup", h.Singup).Methods("POST")
	a.HandleFunc("/signin", h.SignIn).Methods("POST")
	api := r.PathPrefix("/api").Subrouter()
	api.Use(identifyMiddleware)
	api.HandleFunc("/notebooks", h.Create).Methods("POST")
	api.HandleFunc("/notebooks/{notebook_id}", h.Get).Methods("GET")
	api.HandleFunc("/notebooks", h.Update).Methods("PUT")
	api.HandleFunc("/notebooks/{notebook_id}", h.Delete).Methods("DELETE")
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
	var body model.CreateNotebookInput
	ctx := r.Context()
	userRaw := ctx.Value("user")
	user, ok := userRaw.(*model.UserEntityDomain)
	if !ok || user == nil {
		http.Error(w, "user not in context", http.StatusForbidden)
	}
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
func (h *Handler) Singup(w http.ResponseWriter, r *http.Request) {
	var body model.SignUpUserDTO
	ctx := context.TODO()
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "falha ao decodificar corpo da requisicao", http.StatusInternalServerError)
	}
	resp, err := h.service.SignUp(ctx, body)
	if err != nil {
		http.Error(w, "falha ao fazer signup", http.StatusInternalServerError)
	}
	println("DB response:", resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp.Token))
}
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var body model.SignInUserDTO
	ctx := context.TODO()
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "falha ao decodificar corpo da requisicao"+err.Error(), http.StatusInternalServerError)
	}
	resp, err := h.service.SignIn(ctx, body)
	if err != nil {
		http.Error(w, "falha ao fazer signin", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Falha na decodificacao do usuário", http.StatusInternalServerError)
	}
}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["notebook_id"]
	var nb model.Notebook
	ctx := r.Context()
	userRaw := ctx.Value("user")
	user, ok := userRaw.(*model.UserEntityDomain)
	if !ok || user == nil {
		http.Error(w, "user not in context", http.StatusForbidden)
		return
	}
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
	var body model.Notebook
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
