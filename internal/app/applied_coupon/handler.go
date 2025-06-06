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

	if err := h.svc.CreateAppliedCouponService(data.CartID, data.CouponID); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, data)
}

func (h *AppliedCouponHandler) GetAppliedCouponByCartAndCouponID(c *gin.Context) {
	cartID := c.Param("cartID")
	var data []AppliedCoupon

	if err := h.svc.GetAppliedCouponByCartIDService(&data, cartID); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *AppliedCouponHandler) DeleteAppliedCoupon(c *gin.Context) {
	var data AppliedCoupon

	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body is not valid",
		})
	}

	if err := h.svc.DeleteAppliedCouponService(data); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
