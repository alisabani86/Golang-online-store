package presentation

type ShopingCart struct {
	ID        int `json:"id" db:"id"`
	UserID    int `json:"user_id" db:"user_id"`
	ProductID int `json:"product_id" db:"product_id"`
	Quantity  int `json:"quantity" db:"quantity"`
}

type ShopingCartDetail struct {
	ID       int     `json:"cart_id" db:"id"`
	Name     string  `json:"name" db:"product_name"`
	Quantity int     `json:"quantity" db:"quantity"`
	Price    float32 `json:"price" db:"price"`
	Total    int     `json:"total" db:"total"`
}
