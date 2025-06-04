package cart

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CartService interface {
	GetCarts(data *[]Cart) error
	GetCart(data *Cart, id string) error
	CreateCart(data *Cart) error
	FindOrCreateCart(data *Cart, cartID uint) error
	RecalculateTotalPriceService(cartID uint, totalPrice float64) error
	RecalculateTotalDiscountService(cartID uint, totalDiscount float64) error
}

type service struct {
	repo CartRepository
}

func NewCartService(repo CartRepository) CartService {
	return &service{repo: repo}
}

func (svc *service) GetCarts(data *[]Cart) error {
	if err := svc.repo.RepositoryGetCarts(data); err != nil {
		return err
	}
	return nil
}

func (svc *service) GetCart(data *Cart, id string) error {
	if id == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Cart id is required")
	}

	if err := svc.repo.RepositoryGetCart(data, id); err != nil {
		return err
	}
	return nil
}

func (svc *service) CreateCart(data *Cart) error {
	if data == nil {
		return utils.NewDomainError(http.StatusBadRequest, "Request body is empty")
	}

	if data.TotalPrice < 0 || data.DiscountAmount < 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Total price and discount amount must be greater than 0")
	}

	if data.CartStatus != StatusPending {
		return utils.NewDomainError(http.StatusBadRequest, "Cart status must be pending")
	}

	if err := svc.repo.RepositoryCreateCart(data); err != nil {
		return err
	}

	return nil
}

func (svc *service) FindOrCreateCart(data *Cart, cartID uint) error {
	if err := svc.repo.RepositoryFindOrCreateCart(data, cartID); err != nil {
		return err
	}

	return nil
}

func (svc *service) RecalculateTotalPriceService(cartID uint, totalPrice float64) error {
	if cartID <= 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Cart id is required")
	}

	if totalPrice < 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Total price must be greater than 0")
	}

	if err := svc.repo.RepositoryUpdateCartTotalPrice(cartID, totalPrice); err != nil {
		return err
	}

	return nil
}

func (svc *service) RecalculateTotalDiscountService(cartID uint, totalDiscount float64) error {
	if cartID <= 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Cart id is required")
	}

	if totalDiscount < 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Total discount must be greater than 0")
	}

	if err := svc.repo.RepositoryUpdateCartTotalDiscount(cartID, totalDiscount); err != nil {
		return err
	}

	return nil
}
