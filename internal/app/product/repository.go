package product

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

type ProductRepository interface {
	RepositoryGetProducts(data *[]Product) error
	RepositoryGetProduct(data *Product, id string) error
	RepositoryCreatProduct(data *Product) error
}

type repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &repository{db: db}
}

func (r *repository) RepositoryGetProducts(data *[]Product) error {
	if err := r.db.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No product found")
	}

	return nil
}

func (r *repository) RepositoryGetProduct(data *Product, id string) error {
	if err := r.db.First(&data, id).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No product found")
	}

	return nil
}

func (r *repository) RepositoryCreatProduct(data *Product) error {
	if err := r.db.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}
