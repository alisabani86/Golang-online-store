package service

import (
	"context"
	"server/internal/presentation"
)

type Service interface {
	CreateUser(ctx context.Context, req *presentation.CreateUserRequest) (*presentation.CreateUserResponse, error)
	Login(ctx context.Context, req *presentation.LoginUserRequest) (*presentation.LoginUserResponse, error)
	GetProductBasedOnCategory(ctx context.Context, category string) (*[]presentation.Product, error)
	GetProductById(ctx context.Context, id int) (*presentation.Product, error)
	AddCart(ctx context.Context, psc presentation.ShopingCart) error
	GetListCart(ctx context.Context, userid int) ([]presentation.ShopingCartDetail, error)
	DeleteShopingCart(ctx context.Context, cartid int) error
	GetTotalPrice(ctx context.Context, id int) (int, error)
	GetBalance(ctx context.Context, id int) (*presentation.Accounts, error)
	InsertTrancation(ctx context.Context, accountid int, amount int, sofnumber string) error
	UpdateBalance(ctx context.Context, userid int, totalprice int) error
	InsertOrder(ctx context.Context, userid int, total float64, cartid int) error
	UpdateShoppingCart(ctx context.Context, cartid int) error
}
