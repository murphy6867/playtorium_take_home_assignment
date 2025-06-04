package applied_coupon

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/cart_item"
	"github.com/murphy6867/productcheckout/internal/app/coupon"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
	"time"
)

type AppliedCouponService interface {
	CreateAppliedCouponService(cartID uint, couponID uint) error
	GetAppliedCouponByCartIDService(data *[]AppliedCoupon, cartID string) error
	DeleteAppliedCouponService(coupon AppliedCoupon) error
}

type service struct {
	repo            AppliedCouponRepository
	cartService     cart.CartService
	couponService   coupon.CouponService
	cartItemService cart_item.CartItemService
}

func NewAppliedCouponService(
	repo AppliedCouponRepository,
	cartService cart.CartService,
	couponService coupon.CouponService,
	cartItemService cart_item.CartItemService,
) AppliedCouponService {
	return &service{
		repo:            repo,
		cartService:     cartService,
		couponService:   couponService,
		cartItemService: cartItemService,
	}
}

func (s *service) CreateAppliedCouponService(cartID uint, couponID uint) error {
	if err := s.repo.RepoValidateIsExistCouponApplied(cartID, couponID); err != nil {
		return err
	}

	appliedCoupon := AppliedCoupon{
		CartID:    cartID,
		CouponID:  couponID,
		AppliedAt: time.Now(),
	}
	if err := s.repo.RepoCreateAppliedCoupon(&appliedCoupon); err != nil {
		return err
	}

	if err := s.recalculateTotalDiscount(cartID); err != nil {
		if deleteErr := s.repo.RepoDeleteAppliedCoupon(fmt.Sprint(cartID), fmt.Sprint(couponID)); deleteErr != nil {
			return fmt.Errorf("failed to recalculate discount and rollback failed: %v, %v", err, deleteErr)
		}
		return err
	}

	return nil

}

func (s *service) GetAppliedCouponByCartIDService(data *[]AppliedCoupon, cartID string) error {
	if cartID == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Cart ID is required")
	}

	return s.repo.RepoGetAppliedCouponByCartID(data, cartID)
}

func (s *service) DeleteAppliedCouponService(coupon AppliedCoupon) error {
	if fmt.Sprint(coupon.CartID) == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Cart ID is required")
	}

	if fmt.Sprint(coupon.CouponID) == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon ID is required")
	}

	if err := s.repo.RepoDeleteAppliedCoupon(fmt.Sprint(coupon.CartID), fmt.Sprint(coupon.CouponID)); err != nil {
		return err
	}

	return nil
}
