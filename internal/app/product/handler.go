package product

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type ProductHandler struct {
	svc ProductService
}

func NewProductHandler(svc ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	product := make([]Product, 0)
	if err := h.svc.ServiceGetProducts(&product); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	ID := c.Param("id")
	var product Product
	if err := h.svc.ServiceGetProduct(&product, ID); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) PostProduct(c *gin.Context) {
	var body Product

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
		})
		return
	}

	if err := h.svc.ServiceCreateProduct(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, body)
}
