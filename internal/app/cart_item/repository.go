package cart_item

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

type Repository interface {
	RepositoryCreateCartItem(data *CartItem) error
	RepositoryGetCartItemByCartAndProductID(data *CartItem, cartID uint, productID uint) (*CartItem, error)
	RepositoryGetCartItemsByCartID(data *[]CartItem, cartID string) error
	RepositoryUpdateCartItem(data *CartItem, id string) error
	RepoEditCartItem(data *CartItem, quantity int32) error
	RepoGetCartItemByProductCategory(data *[]CartItem, cartID string, categoryName string) error
}

type repository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) RepositoryCreateCartItem(data *CartItem) error {
	if err := r.db.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, fmt.Sprintf("%s", err))
	}

	return nil
}

func (r *repository) RepositoryGetCartItemsByCartID(data *[]CartItem, cartID string) error {
	if err := r.db.
		Preload("Product").
		Where("cart_id", cartID).
		Find(&data).
		Error; err != nil {

		return utils.NewDomainError(http.StatusNotFound, "No cart item found")
	}

	return nil
}

func (r *repository) RepositoryGetCartItemByCartAndProductID(data *CartItem, cartID uint, productID uint) (*CartItem, error) {
	if err := r.db.
		Where("product_id = ? AND cart_id = ?", productID, cartID).
		Find(data).
		Error; err != nil {
		return nil, utils.NewDomainError(http.StatusNotFound, "No cart item found")
	}

	if data.CartID == 0 || data.ProductID == 0 {
		return nil, utils.NewDomainError(http.StatusNotFound, "No cart item found")
	}

	return data, nil
}

func (r *repository) RepositoryUpdateCartItem(data *CartItem, id string) error {
	if err := r.db.Model(&CartItem{}).
		Where("id = ?", data.ID).
		Select("Quantity", "TotalPrice").
		Updates(CartItem{
			Quantity:   data.Quantity,
			TotalPrice: data.TotalPrice,
		}).Error; err != nil {
		return utils.NewDomainError(http.StatusInternalServerError, "Server can not update cart item")
	}

	return nil
}

func (r *repository) RepoEditCartItem(data *CartItem, quantity int32) error {
	if err := r.db.Model(data).
		Where("id = ?", data.ID).
		Update("quantity", quantity).
		Error; err != nil {
		return utils.NewDomainError(http.StatusInternalServerError, "Server can not update cart item")
	}

	return nil
}

func (r *repository) RepoGetCartItemByProductCategory(
	data *[]CartItem,
	cartID string,
	categoryName string) error {
	if err := r.db.
		Preload("Product").
		Joins("JOIN products ON products.id = cart_items.product_id").
		Where("cart_items.cart_id = ? AND products.category_name = ?", cartID, categoryName).
		Find(data).
		Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No cart item found")
	}

	return nil
}
