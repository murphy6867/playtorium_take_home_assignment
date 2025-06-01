package coupon

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func ServiceGetCoupons(data *[]Coupon) error {
	if err := RepositoryGetCoupons(data); err != nil {
		return err
	}
	return nil
}

func ServiceCreateCoupon(data *Coupon) error {
	if data == nil {
		return utils.NewDomainError(http.StatusBadRequest, "Request body is empty")
	}

	if !data.CouponType.IsValidCouponType() {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon type is invalid or missing")
	}

	if !data.CalculateMode.IsValidCalculateMode() {
		return utils.NewDomainError(http.StatusBadRequest, "Calculate mode is invalid or missing")
	}

	if !data.IsValidValue() {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon value is invalid")
	}

	if !data.IsValidCampaignWithCategory() {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon campaign with category is invalid")
	}

	if err := RepositoryCreateCoupon(data); err != nil {
		return err
	}

	return nil
}
