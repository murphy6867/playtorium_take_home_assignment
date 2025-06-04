package applied_coupon

import (
	"fmt"
	"github.com/murphy6867/productcheckout/internal/app/coupon"
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AppliedCouponRepository interface {
	RepoCreateAppliedCoupon(data *AppliedCoupon) error
	RepoGetAppliedCouponByCartAndCouponID(cartID uint, couponID uint) error
	RepoGetAppliedCouponByCartID(data *[]AppliedCoupon, cartID string) error
	RepoDeleteAppliedCoupon(cartID string, couponID string) error
	RepoValidateIsExistCouponApplied(cartID uint, newCouponID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewAppliedCouponRepository(db *gorm.DB) AppliedCouponRepository {
	return &repository{db: db}
}

func (r *repository) RepoCreateAppliedCoupon(data *AppliedCoupon) error {
	if err := r.db.Create(data).Error; err != nil {
		log.Printf("Error: %s", err)
		return utils.NewDomainError(http.StatusNotImplemented, "Coupon is not valid or missing")
	}

	return nil
}

func (r *repository) RepoGetAppliedCouponByCartAndCouponID(cartID, couponID uint) error {
	if err := r.db.
		Where("cart_id = ? AND coupon_id = ?", cartID, couponID).
		First(&AppliedCoupon{}).
		Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No Applied Coupon found")
	}

	return nil
}

func (r *repository) RepoGetAppliedCouponByCartID(data *[]AppliedCoupon, cartID string) error {
	if err := r.db.
		Preload("Cart").
		Preload("Coupon").
		Where("cart_id = ?", cartID).
		Find(data).
		Error; err != nil {
		log.Printf("Error: %s", err)
		return utils.NewDomainError(http.StatusNotFound, "No Applied Coupon found")
	}

	if len(*data) == 0 {
		return utils.NewDomainError(http.StatusNotFound, "No Applied Coupon found")
	}

	return nil
}

func (r *repository) RepoDeleteAppliedCoupon(cartID string, couponID string) error {
	if err := r.db.
		Unscoped().
		Where("cart_id = ? AND coupon_id = ?", cartID, couponID).
		Delete(&AppliedCoupon{}).
		Error; err != nil {
		return utils.NewDomainError(http.StatusInternalServerError, "Server can not delete applied coupon")
	}

	return nil
}

func (r *repository) RepoValidateIsExistCouponApplied(cartID uint, newCouponID uint) error {
	var newCoupon coupon.Coupon
	if err := r.db.First(&newCoupon, newCouponID).Error; err != nil {
		return utils.NewDomainError(
			http.StatusInternalServerError,
			"Error fetching new coupon details",
		)
	}

	var existingAppliedCoupon AppliedCoupon
	err := r.db.Joins("LEFT JOIN coupons ON applied_coupons.coupon_id = coupons.id").
		Where("applied_coupons.cart_id = ? AND coupons.coupon_type = ?", cartID, newCoupon.CouponType).
		First(&existingAppliedCoupon).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	if err == nil {
		return utils.NewDomainError(
			http.StatusBadRequest,
			fmt.Sprintf("Cart already has a coupon of type %s applied", newCoupon.CouponType),
		)
	}

	return utils.NewDomainError(
		http.StatusInternalServerError,
		fmt.Sprintf("Error checking existing coupons: %v", err),
	)
}
