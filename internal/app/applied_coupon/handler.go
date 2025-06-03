package applied_coupon

import "github.com/gin-gonic/gin"

type AppliedCouponHandler struct {
	svc AppliedCouponService
}

func NewAppliedCouponHandler(svc AppliedCouponService) *AppliedCouponHandler {
	return &AppliedCouponHandler{svc: svc}
}

func (h *AppliedCouponHandler) CreateAppliedCoupon(c *gin.Context) {}
