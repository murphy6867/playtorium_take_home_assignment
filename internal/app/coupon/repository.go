package coupon

import (
	"github.com/murphy6867/productcheckout/internal/config"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func RepositoryGetCoupons(data *[]Coupon) error {
	if err := config.DB.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No category found")
	}

	return nil
}

func RepositoryCreateCoupon(data *Coupon) error {
	if !data.IsValidCategory(config.DB) && data.CalculateMode == CalculateModePercentByCategory {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon category is invalid or missing")
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}
