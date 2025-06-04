package cart_item

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/product"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func (c *CartItem) validateInputValue(data *CartItem) bool {
	if fmt.Sprint(data.CartID) == "" {
		return false
	}

	if fmt.Sprint(data.ProductID) == "" {
		return false
	}

	if data.Quantity == 0 {
		return false
	}

	return true
}

func (s *service) validateCartAndProduct(data *CartItem) error {
	var cartInstance cart.Cart
	var productInstance product.Product

	if err := s.cartService.GetCart(&cartInstance, fmt.Sprint(data.CartID)); err != nil {
		return utils.NewDomainError(http.StatusBadRequest, "Invalid cart ID")
	}

	if err := s.productService.ServiceGetProduct(&productInstance, fmt.Sprint(data.ProductID)); err != nil {
		return utils.NewDomainError(http.StatusBadRequest, "Invalid product ID")
	}

	data.Cart = &cartInstance
	data.Product = &productInstance
	return nil
}

func (s *service) recalculateTotalPrice(cartID uint) error {
	var cartItems []CartItem

	if err := s.repo.RepositoryGetCartItemsByCartID(&cartItems, fmt.Sprint(cartID)); err != nil {
		return err
	}

	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += item.TotalPrice
	}

	return s.cartService.RecalculateTotalPriceService(cartID, totalPrice)
}
