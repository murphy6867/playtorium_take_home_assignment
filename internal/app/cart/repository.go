package cart

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

type CartRepository interface {
	RepositoryGetCarts(*[]Cart) error
	RepositoryGetCart(*Cart, string) error
	RepositoryCreateCart(*Cart) error
	RepositoryFindOrCreateCart(data *Cart, cartID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &repository{db: db}
}

func (r *repository) RepositoryGetCarts(data *[]Cart) error {
	if err := r.db.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No cart found")
	}

	return nil
}

func (r *repository) RepositoryGetCart(data *Cart, id string) error {
	if err := r.db.First(&data, id).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No cart found")
	}

	return nil
}

func (r *repository) RepositoryCreateCart(data *Cart) error {
	if err := r.db.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}

func (r *repository) RepositoryFindOrCreateCart(data *Cart, cartID uint) error {
	if err := r.db.Find(&data, cartID).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No cart found")
	}

	return nil
}
