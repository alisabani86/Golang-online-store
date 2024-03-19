package presentation



type Accounts struct {
	ID int `json:"id" db:"id"`
	UserID int `json:"user_id" db:"user_id"`
	AccountNumber string`json:"account_number" db:"account_number"`
	Balance  int `json:"balance" db:"balance"`
}

