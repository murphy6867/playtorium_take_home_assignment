package cart

import (
	"gorm.io/gorm"
)

type StatusEnum string

const (
	StatusPending StatusEnum = "pending"
	StatusSuccess StatusEnum = "success"
	StatusCancel  StatusEnum = "cancel"
)

type Cart struct {
	gorm.Model
	TotalPrice     float64    `json:"total_price"`
	DiscountAmount float64    `json:"discount_amount"`
	Status         StatusEnum `json:"status"`
}
