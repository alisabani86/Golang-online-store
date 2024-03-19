package repository

import (
	"context"
	"online_store/internal/presentation"
)

type Repository interface {
	CreateUser(ctx context.Context, user *presentation.User) (*presentation.User, error)
	GetUserByEmail(ctx context.Context, email string) (*presentation.User, error)
	GetProductBasedOnCategory(ctx context.Context, category string) (*[]presentation.Product, error)
	GetProductById(ctx context.Context, id int) (*presentation.Product, error)
	AddShopingCart(ctx context.Context, psc presentation.ShopingCart) error
	GetShopingCart(ctx context.Context, id int) ([]presentation.ShopingCartDetail, error)
	DeleteShopingCart(ctx context.Context, Shopingcartid int) error
	GetTotalprice(ctx context.Context, id int) (int, error)
	GetBalance(ctx context.Context, id int) (*presentation.Accounts, error)
	InsertTrancation(ctx context.Context, accountid int, amount int, sofnumber string) error
	UpdateBalance(ctx context.Context, accountID int, newBalance int) error
	InsertOrder(ctx context.Context, userID int, total float64, cartID int) error
	UpdateShoppingCart(ctx context.Context, cartid int) error
}
