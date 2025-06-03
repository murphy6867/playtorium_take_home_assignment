package applied_coupon

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AppliedCouponRepository interface {
	RepoCreateAppliedCoupon(data *AppliedCoupon) error
	RepoGetAppliedCouponByCartAndCouponID(data *AppliedCoupon) (*AppliedCoupon, error)
	RepoGetAppliedCouponByCartID(data *[]AppliedCoupon, cartID string) error
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

func (r *repository) RepoGetAppliedCouponByCartAndCouponID(data *AppliedCoupon) (*AppliedCoupon, error) {
	if err := r.db.
		Preload("Coupon").
		Where("cart_id = ? AND coupon_id = ?", data.CartID, data.CouponID).
		First(data).
		Error; err != nil {
		return nil, utils.NewDomainError(http.StatusNotFound, "No Applied Coupon found")
	}

	return data, nil
}

func (r *repository) RepoGetAppliedCouponByCartID(data *[]AppliedCoupon, cartID string) error {
	if err := r.db.
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
