package models

import "time"

type Account struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	AccountName string    `gorm:"not null" json:"account_name"`
	AccountType string    `json:"account_type"`
	Balance     float64   `gorm:"default:0" json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
