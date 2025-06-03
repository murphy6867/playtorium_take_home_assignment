package applied_coupon

import "gorm.io/gorm"

type AppliedCouponRepository interface {
}

type repository struct {
	db *gorm.DB
}

func NewAppliedCouponRepository(db *gorm.DB) AppliedCouponRepository {
	return &repository{db: db}
}
