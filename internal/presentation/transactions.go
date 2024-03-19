package presentation

import "time"

type Transaction struct {
	ID                   int       `json:"id" db:"id"`
	SofNumber            string    `json:"sof_number" db:"sof_number"`
	DofNumber            string    `json:"dof_number" db:"dof_number"`
	AccountID            int       `json:"account_id" db:"account_id"`
	Amount               float64   `json:"amount" db:"amount"`
	TransactionType      int       `json:"transaction_type" db:"transaction_type"`
	TransactionsDatetime time.Time `json:"transactions_datetime" db:"transactions_datetime"`
}
