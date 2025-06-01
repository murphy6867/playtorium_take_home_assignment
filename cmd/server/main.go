package main

import (
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/app/cart"
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
	router := gin.Default()

	router.GET("/category", category.GetCategory)
	router.POST("/category", category.PostCategory)

	router.GET("/product", product.GetProducts)
	router.GET("/product/:id", product.GetProduct)
	router.POST("/product", product.PostProduct)

	router.GET("/coupon", coupon.GetCouponsHandler)
	router.POST("/coupon", coupon.PostCoupon)

	router.GET("/cart", cart.GetCartsHandler)

	//router.POST("/order", handler.OrderCreateHandler)
	//
	//router.POST("/posts", repository.PostsCreate)
	//router.PUT("/posts/:id", repository.PostUpdate)
	//
	//router.GET("/posts", repository.PostsIndex)
	//router.GET("/posts/:id", repository.PostsShow)
	//
	//router.DELETE("/posts/:id", repository.PostDelete)

	router.Run()
}
