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
	RepositoryGetCartItemByCartID(data *[]CartItem, cartID string) error
	RepositoryUpdateCartItem(data *CartItem, id string) error
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

func (r *repository) RepositoryGetCartItemByCartID(data *[]CartItem, cartID string) error {
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
	//if err := r.db.First(&data, id).Error; err != nil {
	//	return nil, utils.NewDomainError(http.StatusNotFound, "No cart item found")
	//}

	//if err := r.db.Model(data).Find(CartItem{CartID: cartID, ProductID: productID}).Error; err != nil {
	//	return nil, utils.NewDomainError(http.StatusNotFound, "No cart item found")
	//}

	if err := r.db.Where("product_id = ? AND cart_id = ?", productID, cartID).Find(data).Error; err != nil {
		return nil, utils.NewDomainError(http.StatusNotFound, "No cart item found")
	}

	if data.CartID == 0 || data.ProductID == 0 {
		return nil, utils.NewDomainError(http.StatusNotFound, "No cart item found")
	}
	fmt.Println("=== Putang ===")

	return data, nil
}

func (r *repository) RepositoryUpdateCartItem(data *CartItem, id string) error {
	fmt.Println("====== 1 > ", data.ID)
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
