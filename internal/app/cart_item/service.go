package cart_item

import (
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/product"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CartItemService interface {
	CreateCartItem(data *CartItem) error
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

func (s *service) CreateCartItem(data *CartItem) error {
	var instanceCart cart.Cart
	var instanceProduct product.Product

	if err := s.cartService.GetCart(&instanceCart, "1"); err != nil {
		return utils.NewDomainError(http.StatusBadRequest, "Invalid cart ID")
	}

	if err := s.productService.ServiceGetProduct(&instanceProduct, "1"); err != nil {
		return utils.NewDomainError(http.StatusBadRequest, "Invalid product ID")
	}

	if err := s.repo.RepositoryCreateCartItem(data); err != nil {
		return err
	}

	return nil
}
