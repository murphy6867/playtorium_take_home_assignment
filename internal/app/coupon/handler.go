package coupon

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CouponHandler struct {
	svc CouponService
}

func NewCouponHandler(svc CouponService) *CouponHandler {
	return &CouponHandler{svc: svc}
}

func (h *CouponHandler) GetCouponsHandler(c *gin.Context) {
	coupon := make([]Coupon, 0)

	if err := h.svc.ServiceGetCoupons(&coupon); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, coupon)
}

func (h *CouponHandler) PostCoupon(c *gin.Context) {
	var body Coupon

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
		})
		return
	}

	if err := h.svc.ServiceCreateCoupon(&body); err != nil {
		utils.HandleError(c, err)
		return
	}
}

func (h *CouponHandler) GetCouponById(c *gin.Context) {
	id := c.Param("id")
	var coupon Coupon

	if err := h.svc.ServiceGetCouponByID(&coupon, id); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, coupon)
}
