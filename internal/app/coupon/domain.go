package coupon

import (
	"github.com/murphy6867/productcheckout/internal/app/category"
	"gorm.io/gorm"
)

type TypeEnum string
type CalculateModeEnum string

const (
	TypeCoupon   TypeEnum = "coupon"
	TypeOnTop    TypeEnum = "on_top"
	TypeSeasonal TypeEnum = "seasonal"
)

const (
	CalculateModeFixed             CalculateModeEnum = "fixed"
	CalculateModePercent           CalculateModeEnum = "percent"
	CalculateModePercentByCategory CalculateModeEnum = "percent_by_category"
	CalculateModePointDiscount     CalculateModeEnum = "point_discount"
	CalculateModeBuyXDiscountY     CalculateModeEnum = "buy_x_discount_y"
)

type Coupon struct {
	gorm.Model
	CouponType      TypeEnum           `json:"coupon_type"`
	CalculateMode   CalculateModeEnum  `json:"calculate_mode"`
	PercentDiscount *float64           `json:"percent_discount"`
	FlatDiscount    *float64           `json:"flat_discount"`
	MinOrderAmount  *float64           `json:"min_order_amount"`
	MaxDiscountCap  *float64           `json:"max_discount_cap"`
	PointUsed       *float64           `json:"point_used"`
	IsActive        bool               `json:"is_active"`
	CategoryName    *string            `json:"category_name"`
	Category        *category.Category `gorm:"foreignKey:CategoryName;references:Name"`
}
