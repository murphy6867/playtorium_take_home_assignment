package category

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func GetCategory(c *gin.Context) {
	category := make([]Category, 0)

	if err := ServiceGetCategory(c, &category); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, category)
}

func PostCategory(c *gin.Context) {
	var body Category

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
		})
		return
	}

	if err := ServiceCreateCategory(c, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

}
