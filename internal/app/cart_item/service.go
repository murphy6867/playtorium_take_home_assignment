package cart_item

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/product"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CartItemService interface {
	CreateCartItemService(data *CartItem) error
	GetCartItemByCartIDService(data *[]CartItem, cartID string) error
}

type service struct {
	repo           Repository
	cartService    cart.CartService
	productService product.ProductService
}

func NewCartItemService(
	repo Repository,
	cartService cart.CartService,
	productService product.ProductService) CartItemService {
	return &service{
		repo:           repo,
		cartService:    cartService,
		productService: productService,
	}
}

func (s *service) CreateCartItemService(data *CartItem) error {
	if fmt.Sprint(data.CartID) == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Cart ID is required")
	}

	if fmt.Sprint(data.ProductID) == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Product ID is required")
	}

	if data.Quantity == 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Quantity is required")
	}

	var instanceCart cart.Cart
	var instanceProduct product.Product

	if err := s.cartService.GetCart(&instanceCart, fmt.Sprint(data.CartID)); err != nil {
		return utils.NewDomainError(http.StatusBadRequest, "Invalid cart ID")
	}

	if err := s.productService.ServiceGetProduct(&instanceProduct, fmt.Sprint(data.ProductID)); err != nil {
		return utils.NewDomainError(http.StatusBadRequest, "Invalid product ID")
	}

	instanceCartItem, err := s.repo.RepositoryGetCartItemByCartAndProductID(&CartItem{}, data.CartID, data.ProductID)
	if err != nil {
		data.TotalPrice = instanceProduct.Price * float64(data.Quantity)

		if err := s.repo.RepositoryCreateCartItem(data); err != nil {
			return err
		}
	} else {
		data.Quantity += instanceCartItem.Quantity
		data.TotalPrice += utils.RoundFloat(instanceProduct.Price*float64(data.Quantity), 2)
		data.ID = instanceCartItem.ID
		fmt.Println("====== 2 > ", data.ID)

		if err := s.repo.RepositoryUpdateCartItem(data, fmt.Sprint(data.ID)); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetCartItemByCartIDService(data *[]CartItem, cartID string) error {
	if cartID == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Cart ID is required")
	}

	if err := s.repo.RepositoryGetCartItemByCartID(data, cartID); err != nil {
		return err
	}

	return nil
}
