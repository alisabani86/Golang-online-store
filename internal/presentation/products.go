package presentation

type Product struct {
	ID       int     `json:"id"  db:"id"`
	Name     string  `json:"product_name" db:"product_name"`
	Category string  `json:"category" db:"category"`
	Price    float64 `json:"price" db:"price"`
}
