package models

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	AccountID       uint      `json:"account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"` // "credit" or "debit"
	Description     string    `json:"description"`
	Date            time.Time `json:"date"`
	CreatedAt       time.Time `json:"created_at"`
}
