package presentation

import "time"

type ShoppingOrder struct {
	ID      int `json:"id" db:"id"`
	OrderID int `json:"order_id" db:"order_id"`
	CartID  int `json:"cart_id" db:"cart_id"`
}
type Order struct {
	ID          int       `json:"id" db:"id"`
	OrderID     int       `json:"order_id" db:"order_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	OrderDate   time.Time `json:"order_date" db:"order_date"`
	TotalAmount float64   `json:"total_amount" db:"total_amount"`
}
