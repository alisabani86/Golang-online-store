package service

import (
	"context"
	"server/internal/presentation"
)

func (p *service) GetProductBasedOnCategory(ctx context.Context, category string) (*[]presentation.Product, error) {

	product, err := p.Repository.GetProductBasedOnCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return product, nil

}

func (p *service) GetProductById(ctx context.Context, id int) (*presentation.Product, error) {

	product, err := p.Repository.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil

}

func (p *service) AddCart(ctx context.Context, psc presentation.ShopingCart) error {

	err := p.Repository.AddShopingCart(ctx, psc)
	if err != nil {
		return err
	}
	return nil
}

func (p *service) GetListCart(ctx context.Context, userid int) ([]presentation.ShopingCartDetail, error) {
	cart, err := p.Repository.GetShopingCart(ctx, userid)
	if err != nil {
		return []presentation.ShopingCartDetail{}, err
	}
	return cart, nil
}

func (p *service) DeleteShopingCart(ctx context.Context, cartid int) error {
	err := p.Repository.DeleteShopingCart(ctx, cartid)
	if err != nil {
		return err
	}
	return nil
}
