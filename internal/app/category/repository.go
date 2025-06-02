package category

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

type Repository interface {
	RepositoryGetCategory(data *Category, id string) error
	RepositoryGetCategories(data *[]Category) error
	RepositoryCreatCategory(data *Category) error
}

type repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) RepositoryGetCategory(data *Category, id string) error {
	if err := r.db.First(&data, id).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No category found")
	}

	return nil
}

func (r *repository) RepositoryGetCategories(data *[]Category) error {
	if err := r.db.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No category found")
	}

	return nil
}

func (r *repository) RepositoryCreatCategory(data *Category) error {
	if err := r.db.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}
