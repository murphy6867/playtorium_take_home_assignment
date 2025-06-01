package applied_coupon

import (
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/coupon"
	"gorm.io/gorm"
	"time"
)

type AppliedCoupon struct {
	gorm.Model
	CartID    uint      `json:"cart_id"`
	CouponID  uint      `json:"coupon_id"`
	AppliedAt time.Time `json:"applied_at"`

	Cart   *cart.Cart     `gorm:"foreignKey:CartID;references:ID"`
	Coupon *coupon.Coupon `gorm:"foreignKey:CouponID;references:ID"`
}
