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
