package applied_coupon

import (
	"github.com/murphy6867/productcheckout/internal/app/cart"
	"github.com/murphy6867/productcheckout/internal/app/coupon"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
	"time"
)

type AppliedCouponService interface {
	CreateAppliedCouponService(data *AppliedCoupon) error
}

type service struct {
	repo          AppliedCouponRepository
	cartService   cart.CartService
	couponService coupon.CouponService
}

func NewAppliedCouponService(
	repo AppliedCouponRepository,
	cartService cart.CartService,
	couponService coupon.CouponService,
) AppliedCouponService {
	return &service{
		repo:          repo,
		cartService:   cartService,
		couponService: couponService,
	}
}

func (s *service) CreateAppliedCouponService(data *AppliedCoupon) error {
	if !s.ValidateCart(data.CartID) {
		return utils.NewDomainError(http.StatusBadRequest, "Cart is not valid")
	}

	if !s.ValidateCoupon(data.CouponID) {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon is not valid")
	}

	data.AppliedAt = time.Now().UTC()

	if s.ValidateExistCouponApplied(data) {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon is already applied")
	}

	if err := s.repo.RepoCreateAppliedCoupon(data); err != nil {
		return err
	}

	return nil
}
