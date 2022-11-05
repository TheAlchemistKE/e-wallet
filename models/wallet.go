package models

import (
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	gorm.Model
	PhoneNumber int       `json:"phone_number" gorm:"primaryKey"`
	Balance     float64   `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}

func (w *Wallet) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["phone_number"] = w.PhoneNumber
	response["balance"] = w.Balance

	return response
}
