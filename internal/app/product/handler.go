package product

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func GetProducts(c *gin.Context) {
	product := make([]Product, 0)
	if err := ServiceGetProducts(c, &product); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, product)
}

func GetProduct(c *gin.Context) {
	ID := c.Param("id")
	var product Product
	if err := ServiceGetProduct(c, &product, ID); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, product)
}

func PostProduct(c *gin.Context) {
	var body Product

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
		})
		return
	}

	if err := ServiceCreateProduct(c, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
}
