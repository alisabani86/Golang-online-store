package service

import (
	"online_store/internal/middleware"
	"online_store/internal/repository"
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
