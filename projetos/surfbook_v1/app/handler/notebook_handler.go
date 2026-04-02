package handler

import "github.com/YagoSchramm/intensivo-surfbook_v1/service"

type NotebookHandler struct {
	srv *service.NotebookService
}

func NewNotebookHandler(srv *service.NotebookService) *NotebookHandler {
	return &NotebookHandler{srv: srv}
}
