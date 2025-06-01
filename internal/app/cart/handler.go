package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
)

func GetCartsHandler(c *gin.Context) {
	if err := ServiceGetCarts(); err != nil {
		utils.HandleError(c, err)
		return
	}
}
