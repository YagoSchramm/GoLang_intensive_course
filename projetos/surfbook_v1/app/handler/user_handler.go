package handler

import "github.com/YagoSchramm/intensivo-surfbook_v1/service"

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
