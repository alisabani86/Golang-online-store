package service

import (
	"server/internal/middleware"
	"server/internal/repository"
	"time"
)

type service struct {
	repository.Repository
	timout time.Duration
	middleware.Middleware
}

func NewService(repository repository.Repository, middleware middleware.Middleware) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
		middleware,
	}
}
