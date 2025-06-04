package applied_coupon

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/cart_item"
	"github.com/murphy6867/productcheckout/internal/app/coupon"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
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

func (s *service) ValidateIsExistCouponApplied(data *AppliedCoupon) bool {

	var existData *[]AppliedCoupon

	if err := s.GetAppliedCouponByCartIDService(existData, fmt.Sprint(data.CartID)); err != nil {
		fmt.Println("===== err: ", err)
		return false
	}

	//type AppliedCoupon struct {
	//	gorm.Model
	//	CartID    uint           `json:"cart_id"`
	//	CouponID  uint           `json:"coupon_id"`
	//	AppliedAt time.Time      `json:"applied_at"`
	//	Cart      *cart.Cart     `gorm:"foreignKey:CartID;references:ID"`
	//	Coupon    *coupon.Coupon `gorm:"foreignKey:CouponID;references:ID"`
	//}

	//for _, eachCoupon := range *data {}

	return true
}

func (s *service) recalculateTotalDiscount(cartID uint) error {
	var appliedCoupons []AppliedCoupon

	if err := s.GetAppliedCouponByCartIDService(&appliedCoupons, fmt.Sprint(cartID)); err != nil {
		return err
	}

	if len(appliedCoupons) == 0 {
		return utils.NewDomainError(http.StatusNotFound, "No coupons applied to this cart")
	}

	if appliedCoupons[0].Cart.PriceAfterDiscount == 0.0 {

	}

	cart := appliedCoupons[0].Cart
	totalDiscount := 0.0

	for _, eachCoupon := range appliedCoupons {
		switch eachCoupon.Coupon.CouponType {
		case "coupon":
			if err := s.applyCouponDiscount(eachCoupon.Coupon, &totalDiscount, cart); err != nil {
				return err
			}
		case "on_top":
			if err := s.applyOnTopDiscount(eachCoupon.Coupon, &totalDiscount, cart); err != nil {
				return err
			}
		case "seasonal":
			if err := s.applySeasonalDiscount(eachCoupon.Coupon, &totalDiscount, cart); err != nil {
				return err
			}
		default:
			return utils.NewDomainError(http.StatusBadRequest, "Unknown coupon type")
		}
	}

	if err := s.cartService.RecalculateTotalDiscountService(cartID, totalDiscount); err != nil {
		return err
	}

	return nil
}

func (s *service) applyCouponDiscount(coupon *coupon.Coupon, totalDiscount *float64, cart *cart.Cart) error {
	switch coupon.CalculateMode {
	case "fixed":
		fmt.Println("======= flat used ===========")

		if cart.TotalPrice <= *coupon.FlatDiscount {
			*totalDiscount += 0.0
		} else {
			*totalDiscount += *coupon.FlatDiscount
		}
		return nil

	case "percent":
		fmt.Println("======= percent used ===========")

		discount := cart.TotalPrice * (*coupon.PercentDiscount / 100.0)
		*totalDiscount += discount
		return nil

	default:
		return utils.NewDomainError(http.StatusBadRequest, "Invalid coupon calculation mode")
	}
}

func (s *service) applyOnTopDiscount(coupon *coupon.Coupon, totalDiscount *float64, cart *cart.Cart) error {
	switch coupon.CalculateMode {
	case "percent_by_category":
		fmt.Println("======= percent by category used ===========")

		var cartItem []cart_item.CartItem
		if err := s.cartItemService.GetCartItemByProductCategory(
			&cartItem,
			fmt.Sprint(cart.ID),
			*coupon.CategoryName); err != nil {
			return err
		}

		for _, item := range cartItem {
			if *coupon.CategoryName == item.Product.CategoryName {
				*totalDiscount += item.TotalPrice * (*coupon.PercentDiscount / 100.0)
			}
		}

		return nil
	case "point_discount":
		fmt.Println("======= point discount used ===========")

		if coupon.PointUsed == nil || coupon.PercentDiscount == nil {
			return utils.NewDomainError(http.StatusBadRequest, "Missing discount data")
		}

		maxDiscount := cart.TotalPrice * (*coupon.PercentDiscount / 100.0)

		if *coupon.PointUsed > maxDiscount {
			return utils.NewDomainError(http.StatusBadRequest, "Invalid point discount")
		}

		*totalDiscount += *coupon.PointUsed

		return nil
	default:
		return utils.NewDomainError(http.StatusBadRequest, "Invalid coupon calculation mode")
	}
}

func (s *service) applySeasonalDiscount(coupon *coupon.Coupon, totalDiscount *float64, cart *cart.Cart) error {
	switch coupon.CalculateMode {
	case "buy_x_discount_y":
		if cart.TotalPrice <= *coupon.FlatDiscount {
			*totalDiscount += 0.0
		} else {
			*totalDiscount += *coupon.FlatDiscount
		}
		return nil
	default:
		return utils.NewDomainError(http.StatusBadRequest, "Invalid coupon calculation mode")
	}
}
