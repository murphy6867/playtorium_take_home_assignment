package applied_coupon

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type AppliedCouponHandler struct {
	svc AppliedCouponService
}

func NewAppliedCouponHandler(svc AppliedCouponService) *AppliedCouponHandler {
	return &AppliedCouponHandler{svc: svc}
}

func (h *AppliedCouponHandler) CreateAppliedCoupon(c *gin.Context) {
	var data AppliedCoupon

	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body is not valid",
		})

		return
	}

	if err := h.svc.CreateAppliedCouponService(&data); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, data)
}
