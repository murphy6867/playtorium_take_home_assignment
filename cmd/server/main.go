package main

import (
	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	categoryRepo := category.NewCategoryRepository(config.DB)
	categorySvc := category.NewCategoryService(categoryRepo)
	categoryHdl := category.NewCategoryHandler(categorySvc)

	couponRepo := coupon.NewCouponRepository(config.DB)
	couponSvc := coupon.NewCouponService(couponRepo)
	couponHdl := coupon.NewCouponHandler(couponSvc)

	productRepo := product.NewProductRepository(config.DB)
	productSvc := product.NewProductService(productRepo)
	productHdl := product.NewProductHandler(productSvc)

	cartRepo := cart.NewCartRepository(config.DB)
	cartSvc := cart.NewCartService(cartRepo)
	cartHdl := cart.NewCartHandler(cartSvc)

	cartItemRepo := cart_item.NewCartItemRepository(config.DB)
	cartItemSvc := cart_item.NewCartItemService(cartItemRepo, cartSvc, productSvc)
	cartItemHdl := cart_item.NewCartItemHandler(cartItemSvc)

	appliedCouponRepo := applied_coupon.NewAppliedCouponRepository(config.DB)
	appliedCouponSvc := applied_coupon.NewAppliedCouponService(appliedCouponRepo, cartSvc, couponSvc)
	appliedCouponHdl := applied_coupon.NewAppliedCouponHandler(appliedCouponSvc)

	router.GET("/category", categoryHdl.GetCategory)
	router.GET("/category/:id", categoryHdl.GetCategoryById)
	router.POST("/category", categoryHdl.PostCategory)

	router.GET("/product", productHdl.GetProducts)
	router.GET("/product/:id", productHdl.GetProduct)
	router.POST("/product", productHdl.PostProduct)

	router.GET("/coupon", couponHdl.GetCouponsHandler)
	router.GET("/coupon/:id", couponHdl.GetCouponById)
	router.POST("/coupon", couponHdl.PostCoupon)

	router.GET("/cart", cartHdl.GetCartsHandler)
	router.GET("/cart/:id", cartHdl.GetCartByIdHandler)
	router.POST("/cart", cartHdl.PostCartHandler)

	router.GET("/cart_item/:cartID", cartItemHdl.GetCartItemByCartIDHandler)
	router.POST("/cart_item", cartItemHdl.CreateCartItemsHandler)

	router.POST("/applied_coupon", appliedCouponHdl.CreateAppliedCoupon)

	router.Run()
}
