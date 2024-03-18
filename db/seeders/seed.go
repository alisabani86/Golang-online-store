package main

import (
	"context"
	"fmt"
	"log"
	"server/db"
)

func seedProducts(ctx context.Context, db *db.Database) error {
	products := []struct {
		ID          int
		ProductName string
		Category    string
		Price       float64
		Quantity    int
	}{
		{1, "Phone X", "Electronics", 699.99, 5},
		{2, "Phone Y", "Electronics", 799.99, 6},
		{3, "Phone Z", "Electronics", 899.99, 10},
		{4, "Tablet A", "Electronics", 599.99, 7},
		{5, "Tablet B", "Electronics", 699.99, 8},
		{6, "Tablet C", "Electronics", 799.99, 5},
		{7, "Tablet D", "Electronics", 899.99, 5},
		{8, "Shirts A", "Non-Electronics", 245.88, 5},
		{9, "Shirts B", "Non-Electronics", 245.88, 5},
		{10, "Shirts C", "Non-Electronics", 245.88, 5},
		{11, "Shirts D", "Non-Electronics", 245.88, 5},
		{12, "Shirts E", "Non-Electronics", 245.88, 5},
		{13, "Shirts F", "Non-Electronics", 245.88, 5},
	}

	query := `
		INSERT INTO products (id, product_name, category, price, quantity)
		VALUES ($1, $2, $3, $4,$5)
	`

	for _, p := range products {
		_, err := db.GetDB().QueryContext(ctx, query, p.ID, p.ProductName, p.Category, p.Price, p.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	db, err := db.NewDatabase()
	ctx := context.Background()

	err = seedProducts(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seed data inserted successfully!")
}
