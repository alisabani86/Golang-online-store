package repository

import (
	"context"
	"log"
	"math/rand"
	"server/internal/presentation"
	"time"
)

func (r repository) GetTotalprice(ctx context.Context, id int) (int, error) {
	query := "SELECT c.id, p.product_name, c.quantity, p.price FROM products p LEFT JOIN shoping_cart c ON p.id = c.product_id WHERE c.user_id = $1 AND c.deleted_at IS NULL;"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []presentation.ShopingCartDetail
	for rows.Next() {
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

	// Calculate the total price
	var totalPrice int
	for _, item := range products {
		totalPrice += item.Total
	}

	// Print the total price (optional)
	log.Printf("Total price for user %d: %d", id, totalPrice)

	return totalPrice, nil
}

func (r repository) GetBalance(ctx context.Context, id int) (*presentation.Accounts, error) {

	u := presentation.Accounts{}

	query := "SELECT id, account_number, balance FROM accounts WHERE user_id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.AccountNumber, &u.Balance)
	if err != nil {
		return &presentation.Accounts{}, err
	}
	return &u, nil

}

func (r repository) InsertTrancation(ctx context.Context, accountid int, amount int, sofnumber string) error {

	query := "INSERT INTO transactions (id, account_id, amount, sof_number) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, rand.Intn(100), accountid, amount, sofnumber, 1, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r repository) UpdateBalance(ctx context.Context, accountID int, newBalance int) error {
	query := "UPDATE accounts SET balance = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, newBalance, accountID)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) InsertOrder(ctx context.Context, userID int, total float64, cartID int) error {
	// Insert into the orders table
	orderQuery := "INSERT INTO orders (user_id, order_date, total_amount) VALUES ($1, $2, $3) RETURNING id"
	var orderID int
	if err := r.db.QueryRowContext(ctx, orderQuery, userID, time.Now(), total).Scan(&orderID); err != nil {
		return err
	}

	// Insert into the shoping_order table
	orderCartQuery := "INSERT INTO shoping_order (id, order_id, cart_id) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, orderCartQuery, rand.Intn(100), orderID, cartID)
	if err != nil {
		return err
	}
	return nil

}

func (r repository) UpdateShoppingCart(ctx context.Context, cartid int) error {
	updateCartQuery := "UPDATE shopping_cart SET deleted_at = NOW() WHERE user_id = $1"
	_, err := r.db.ExecContext(ctx, updateCartQuery, cartid)
	if err != nil {
		return err
	}
	return nil
}
