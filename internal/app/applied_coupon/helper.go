package applied_coupon

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/coupon"
)

func (s *service) ValidateCart(cartID uint) bool {
	if err := s.cartService.GetCart(&cart.Cart{}, fmt.Sprint(cartID)); err != nil {
		return false
	}

	return true
}

func (s *service) ValidateCoupon(couponID uint) bool {
	if err := s.couponService.ServiceGetCouponByID(&coupon.Coupon{}, fmt.Sprint(couponID)); err != nil {
		return false
	}

	return true
}

func (s *service) ValidateExistCouponApplied(data *AppliedCoupon) bool {
	var _, err = s.repo.RepoGetAppliedCouponByCartAndCouponID(data)
	if err != nil {
		return false
	}

	//jsonData, err := json.MarshalIndent(appliedCoupon, "", "    ")
	//if err != nil {
	//	_ = fmt.Errorf("failed to marshal applied coupon data: %w", err)
	//}
	//
	//log.Printf("======> %s", string(jsonData))

	return true
}
