package router

import (
	"online_store/internal/middleware"
	"online_store/internal/service"
)

type Handler struct {
	service.Service
	middleware.Middleware
}

func NewHandler(s service.Service, m middleware.Middleware) *Handler {
	return &Handler{
		Service:    s,
		Middleware: m,
	}
}
