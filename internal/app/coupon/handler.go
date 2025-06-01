package coupon

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func GetCouponsHandler(c *gin.Context) {
	coupon := make([]Coupon, 0)

	if err := ServiceGetCoupons(&coupon); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, coupon)
}

func PostCoupon(c *gin.Context) {
	var body Coupon

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
		})
		return
	}

	if err := ServiceCreateCoupon(&body); err != nil {
		utils.HandleError(c, err)
		return
	}
}
