package product

import (
	"github.com/murphy6867/productcheckout/internal/app/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`

	CategoryName string             `json:"category_name"`
	Category     *category.Category `gorm:"foreignKey:CategoryName;references:Name"`
}
