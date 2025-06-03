package applied_coupon

type AppliedCouponService interface {
}

type service struct {
	repo AppliedCouponRepository
}

func NewAppliedCouponService(repo AppliedCouponRepository) AppliedCouponService {
	return &service{repo: repo}
}
