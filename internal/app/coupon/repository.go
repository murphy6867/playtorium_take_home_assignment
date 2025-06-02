package coupon

import (
	"github.com/murphy6867/productcheckout/internal/config"
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

type Repository interface {
	RepositoryGetCoupons(data *[]Coupon) error
	RepositoryCreateCoupon(data *Coupon) error
}

type repository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) RepositoryGetCoupons(data *[]Coupon) error {
	if err := config.DB.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No category found")
	}

	return nil
}

func (r *repository) RepositoryCreateCoupon(data *Coupon) error {
	if !data.IsValidCategory(config.DB) && data.CalculateMode == CalculateModePercentByCategory {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon category is invalid or missing")
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}
