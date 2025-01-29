package models

import "time"

type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"Name"`
	Balance   float64   `json:"balance"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Transaction struct {
	AccountID       int       `json:"account_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Timestamp       time.Time `json:"timestamp"`
}
