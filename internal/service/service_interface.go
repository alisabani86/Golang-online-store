package service

import (
	"context"
	"server/internal/presentation"
)

type Service interface {
	CreateUser(ctx context.Context, req *presentation.CreateUserRequest) (*presentation.CreateUserResponse, error)
	Login(ctx context.Context, req *presentation.LoginUserRequest) (*presentation.LoginUserResponse, error)
	GetProductBasedOnCategory(ctx context.Context, category string) (*[]presentation.Product, error)
}
