package handler

import "github.com/YagoSchramm/base-auth-v2-session/service"

type UserHandler struct {
	srv *service.UserService
}

func New(srv *service.UserService) *UserHandler {
	return &UserHandler{srv: srv}
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
