package cart_item

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

type Repository interface {
	RepositoryCreateCartItem(data *CartItem) error
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
