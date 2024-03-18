package repository

import (
	"context"
	"server/internal/presentation"
)

type Repository interface {
	CreateUser(ctx context.Context, user *presentation.User) (*presentation.User, error)
	GetUserByEmail(ctx context.Context, email string) (*presentation.User, error)
	GetProductBasedOnCategory(ctx context.Context, category string) (*[]presentation.Product, error)
	GetProductById(ctx context.Context, id int) (*presentation.Product, error)
	AddShopingCart(ctx context.Context, psc presentation.ShopingCart) error
	GetShopingCart(ctx context.Context, id int) ([]presentation.ShopingCartDetail, error)
	DeleteShopingCart(ctx context.Context, Shopingcartid int) error
}
