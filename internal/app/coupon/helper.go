package coupon

import (
	"github.com/murphy6867/productcheckout/internal/app/category"
	"gorm.io/gorm"
)

func (t TypeEnum) IsValidCouponType() bool {
	switch t {
	case TypeCoupon, TypeOnTop, TypeSeasonal:
		return true
	default:
		return false
	}
}

func (t CalculateModeEnum) IsValidCalculateMode() bool {
	switch t {
	case CalculateModeFixed,
		CalculateModePercent,
		CalculateModePercentByCategory,
		CalculateModePointDiscount,
		CalculateModeBuyXDiscountY:
		return true
	default:
		return false
	}
}

func (t *Coupon) IsValidValue() bool {
	if t.FlatDiscount != nil && *t.FlatDiscount <= 0.0 {
		return false
	}

	if t.PercentDiscount != nil && (*t.PercentDiscount <= 0.0 || *t.PercentDiscount > 100.0) {
		return false
	}

	if t.MinOrderAmount != nil && *t.MinOrderAmount < 0.0 {
		return false
	}

	if t.MaxDiscountCap != nil && *t.MaxDiscountCap <= 0.0 {
		return false
	}

	if t.PointUsed != nil && *t.PointUsed < 0.0 {
		return false
	}

	return true
}

func (t *Coupon) IsValidCampaignWithCategory() bool {
	switch t.CouponType {
	case TypeCoupon:
		return t.CalculateMode == CalculateModeFixed || t.CalculateMode == CalculateModePercent
	case TypeOnTop:
		return t.CalculateMode == CalculateModePercentByCategory || t.CalculateMode == CalculateModePointDiscount
	case TypeSeasonal:
		return t.CalculateMode == CalculateModeBuyXDiscountY
	default:
		return false
	}
}

func (t *Coupon) IsValidCategory(db *gorm.DB) bool {
	if t.CategoryName != nil {
		return false
	}
	var cat category.Category
	err := db.Where("name = ?", t.CategoryName).First(&cat).Error
	return err == nil
}
