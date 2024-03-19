package repository

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"online_store/internal/presentation"
	"strings"
)

func (r *repository) GetProductBasedOnCategory(ctx context.Context, category string) (*[]presentation.Product, error) {
	category = strings.ToLower(category)

	query := "SELECT id, product_name, category, price, quantity FROM products WHERE LOWER(category) = $1"
	rows, err := r.db.QueryContext(ctx, query, category)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var products []presentation.Product

	for rows.Next() {
		// Process each row
		var p presentation.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Quantity); err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
		fmt.Printf("Product: ID=%d, Name=%s, Category=%s, Price=%.2f\n", p.ID, p.Name, p.Category, p.Price)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &products, nil

}

func (r repository) GetProductById(ctx context.Context, id int) (*presentation.Product, error) {

	query := "SELECT id, product_name, category, price, quantity FROM products WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	var p presentation.Product
	err := row.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	return &p, nil
}

func (r repository) AddShopingCart(ctx context.Context, psc presentation.ShopingCart) error {
	var lastInsertId int

	id := rand.Intn(10)

	query := "INSERT INTO shoping_cart (id, user_id, product_id, quantity) VALUES ($1, $2, $3,$4) returning id"
	err := r.db.QueryRowContext(ctx, query, id, psc.UserID, psc.ProductID, psc.Quantity).Scan(&lastInsertId)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	psc.ID = int(lastInsertId)
	return nil
}

func (r repository) GetShopingCart(ctx context.Context, id int) ([]presentation.ShopingCartDetail, error) {

	query := "SELECT c.id, p.product_name, c.quantity, p.price FROM products p LEFT JOIN shoping_cart c on p.id = c.product_id WHERE c.user_id = $1 AND c.deleted_at IS NULL;"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var products []presentation.ShopingCartDetail
	for rows.Next() {
		// Process each row
		var p presentation.ShopingCartDetail
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price); err != nil {
			log.Fatal(err)
		}
		p.Total = p.Quantity * int(p.Price)
		products = append(products, p)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return products, nil
}

func (r repository) DeleteShopingCart(ctx context.Context, Shopingcartid int) error {
	query := ` UPDATE shoping_cart SET deleted_at = NOW()
	WHERE id = $1;`
	_, err := r.db.ExecContext(ctx, query, Shopingcartid)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
