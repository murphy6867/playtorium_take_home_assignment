package cart_item

import (
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/product"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`
	CartID     uint    `json:"cart_id"`
	Quantity   int32   `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	
	Cart    *cart.Cart       `gorm:"foreignKey:CartID;references:ID"`
	Product *product.Product `gorm:"foreignKey:ProductID;references:ID"`
}
