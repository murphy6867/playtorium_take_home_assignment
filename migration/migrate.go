package main

import (
	"github.com/murphy6867/productcheckout/internal/app/applied_coupon"
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/cart_item"
	"github.com/murphy6867/productcheckout/internal/app/category"
	"github.com/murphy6867/productcheckout/internal/app/coupon"
	"github.com/murphy6867/productcheckout/internal/app/product"
	"github.com/murphy6867/productcheckout/internal/config"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(
		&category.Category{},
		&product.Product{},
		&coupon.Coupon{},
		&cart.Cart{},
		&applied_coupon.AppliedCoupon{},
		&cart_item.CartItem{},
	)
}
