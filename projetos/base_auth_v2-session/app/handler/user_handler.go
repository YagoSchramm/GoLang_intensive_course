package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/YagoSchramm/base-auth-v2-session/model"
	"github.com/YagoSchramm/base-auth-v2-session/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	srv       *service.UserService
	jwtSecret string
}

type signupResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type signinResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func NewUserHandler(srv *service.UserService, jwtSecret string) *UserHandler {
	return &UserHandler{srv: srv, jwtSecret: jwtSecret}
}

func (h *UserHandler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/signup", h.signup).Methods("POST")
	r.HandleFunc("/signin", h.signin).Methods("POST")
}

func (h *UserHandler) signup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var input model.SignUpUserDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	user, err := h.srv.Create(r.Context(), input)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Erro ao criar usuário")
		return
	}

	h.respondJSON(w, http.StatusCreated, signupResponse{ID: user.ID, Email: user.Email})
}

func (h *UserHandler) signin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var input model.SignInUserDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	user, err := h.srv.Authenticate(r.Context(), input)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Erro ao autenticar")
		return
	}
	if user == nil {
		h.respondError(w, http.StatusUnauthorized, "Credenciais inválidas")
		return
	}

	token, expiresAt, err := h.generateToken(user.ID)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "Erro ao gerar token")
		return
	}

	h.respondJSON(w, http.StatusOK, signinResponse{Token: token, ExpiresAt: expiresAt})
}

func (h *UserHandler) generateToken(userID string) (string, time.Time, error) {
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, expiresAt, nil
}

func (h *UserHandler) respondJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func (h *UserHandler) respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorResponse{Error: message})
}
