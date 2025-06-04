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
	GetCartItemByProductCategory(data *[]CartItem, cartID string, categoryName string) error
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
	if !data.validateInputValue(data) {
		return utils.NewDomainError(http.StatusBadRequest, "Invalid input value")
	}

	if err := s.validateCartAndProduct(data); err != nil {
		return err
	}

	existingItem, err := s.repo.RepositoryGetCartItemByCartAndProductID(&CartItem{}, data.CartID, data.ProductID)
	if err != nil {
		fmt.Println("===== Flow 1: ")
		data.TotalPrice = data.Product.Price * float64(data.Quantity)
		data.TotalPrice = utils.RoundFloat(data.Product.Price*float64(data.Quantity), 2)

		if err := s.repo.RepositoryCreateCartItem(data); err != nil {
			return err
		}

		return s.recalculateTotalPrice(data.CartID)
	}
	fmt.Println("===== Flow 2: ")

	data.Quantity += existingItem.Quantity
	data.TotalPrice = utils.RoundFloat(data.Product.Price*float64(data.Quantity), 2)
	data.ID = existingItem.ID

	if err := s.repo.RepositoryUpdateCartItem(data, fmt.Sprint(data.ID)); err != nil {
		return err
	}

	return s.recalculateTotalPrice(data.CartID)
}

func (s *service) GetCartItemByCartIDService(data *[]CartItem, cartID string) error {
	if cartID == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Cart ID is required")
	}

	return s.repo.RepositoryGetCartItemsByCartID(data, cartID)
}

func (s *service) GetCartItemByProductCategory(data *[]CartItem, cartID string, categoryName string) error {
	if cartID == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Cart ID is required")
	}

	if err := s.repo.RepoGetCartItemByProductCategory(data, cartID, categoryName); err != nil {
		return err
	}

	return nil
}
