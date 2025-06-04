package coupon

import (
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CouponService interface {
	ServiceGetCoupons(data *[]Coupon) error
	ServiceCreateCoupon(data *Coupon) error
	ServiceGetCouponByID(data *Coupon, id string) error
	RepoUpdatePointUse(data *Coupon) error
	ServiceGetCouponByCouponType(data *Coupon, couponType string) error
}

type service struct {
	repo Repository
}

func NewCouponService(repo Repository) CouponService {
	return &service{repo: repo}
}

func (s *service) ServiceGetCoupons(data *[]Coupon) error {
	if err := s.repo.RepositoryGetCoupons(data); err != nil {
		return err
	}
	return nil
}

func (s *service) ServiceCreateCoupon(data *Coupon) error {
	if data == nil {
		return utils.NewDomainError(http.StatusBadRequest, "Request body is empty")
	}

	if !data.CouponType.IsValidCouponType() {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon type is invalid or missing")
	}

	if !data.CalculateMode.IsValidCalculateMode() {
		return utils.NewDomainError(http.StatusBadRequest, "Calculate mode is invalid or missing")
	}

	if !data.IsValidValue() {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon value is invalid")
	}

	if !data.IsValidCampaignWithCategory() {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon campaign with category is invalid")
	}

	if err := s.repo.RepositoryCreateCoupon(data); err != nil {
		return err
	}

	return nil
}

func (s *service) ServiceGetCouponByID(data *Coupon, id string) error {
	if id == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon id is required")
	}

	if err := s.repo.RepositoryGetCouponByID(data, id); err != nil {
		return err
	}

	return nil
}

func (s *service) RepoUpdatePointUse(data *Coupon) error {
	if *data.PointUsed < 0.0 {
		return utils.NewDomainError(http.StatusBadRequest, "Point used is invalid")
	}

	return s.repo.RepoUpdatePointUse(data)
}

func (s *service) ServiceGetCouponByCouponType(data *Coupon, couponType string) error {
	if couponType == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Coupon type is required")
	}

	if err := s.ServiceGetCouponByCouponType(data, couponType); err != nil {
		return err
	}

	return nil
}
