package repository

import (
	"context"
	"fmt"
	"log"
	"server/internal/presentation"
	"strings"
)

func (r *repository) GetProductBasedOnCategory(ctx context.Context, category string) (*[]presentation.Product, error) {
	category = strings.ToLower(category)

	query := "SELECT id, product_name, category, price FROM products WHERE LOWER(category) = $1"
	rows, err := r.db.QueryContext(ctx, query, category)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var products []presentation.Product

	for rows.Next() {
		// Process each row
		var p presentation.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price); err != nil {
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
