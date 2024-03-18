package service

import (
	"context"
	"server/internal/presentation"
)

func (t *service) GetTotalPrice(ctx context.Context, id int) (int, error) {
	return t.Repository.GetTotalprice(ctx, id)
}

func (t *service) GetBalance(ctx context.Context, id int) (*presentation.Accounts, error) {
	return t.Repository.GetBalance(ctx, id)
}

func (t *service) InsertTrancation(ctx context.Context, accountid int, amount int, sofnumber string) error {
	return t.Repository.InsertTrancation(ctx, accountid, amount, sofnumber)
}

func (t *service) UpdateBalance(ctx context.Context, userid int, totalprice int) error {
	return t.Repository.UpdateBalance(ctx, userid, -totalprice)
}

func (t *service) InsertOrder(ctx context.Context, userid int, total float64, cartid int) error {
	return t.Repository.InsertOrder(ctx, userid, total, cartid)
}

func (t *service) UpdateShoppingCart(ctx context.Context, cartid int) error {
	err := t.Repository.UpdateShoppingCart(ctx, cartid)
	if err != nil {
		return err
	}

	return nil
}
