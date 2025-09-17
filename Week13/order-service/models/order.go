package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint   `json:"user_id"`    // from Auth Service
	ProductID uint   `json:"product_id"` // from Product Service
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"` // e.g., "pending", "confirmed", "cancelled"
}
