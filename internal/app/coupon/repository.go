package coupon

import (
	"github.com/murphy6867/productcheckout/internal/config"
	"github.com/murphy6867/productcheckout/internal/utils"
	"gorm.io/gorm"
	"net/http"
)

type Repository interface {
	RepositoryGetCoupons(data *[]Coupon) error
	RepositoryCreateCoupon(data *Coupon) error
	RepositoryGetCouponByID(data *Coupon, id string) error
	RepoUpdatePointUse(data *Coupon) error
	RepoGetCouponByCouponType(data *Coupon, couponType string) error
}

type repository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) RepositoryGetCoupons(data *[]Coupon) error {
	if err := config.DB.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No category found")
	}

	return nil
}

func (r *repository) RepositoryCreateCoupon(data *Coupon) error {
	if !data.IsValidCategory(config.DB) && data.CalculateMode == CalculateModePercentByCategory {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon category is invalid or missing")
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}

func (r *repository) RepositoryGetCouponByID(data *Coupon, id string) error {
	if err := config.DB.First(&data, id).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No coupon found")
	}

	return nil
}
func (r *repository) RepoUpdatePointUse(data *Coupon) error {
	if err := config.DB.Model(data).
		Where("id = ?", data.ID).
		Update("point_used", data.PointUsed).
		Error; err != nil {
		return utils.NewDomainError(http.StatusInternalServerError, "Server can not update coupon item")
	}

	return nil
}

func (r *repository) RepoGetCouponByCouponType(data *Coupon, couponType string) error {
	if err := config.DB.
		Where("coupon_type = ?", couponType).
		Find(&data).
		Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No coupon found")
	}
	return nil
}
