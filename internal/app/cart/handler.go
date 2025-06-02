package cart

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CartHandler struct {
	svc CartService
}

func NewCartHandler(svc CartService) *CartHandler {
	return &CartHandler{svc: svc}
}

func (h *CartHandler) GetCartsHandler(c *gin.Context) {
	data := make([]Cart, 0)
	if err := h.svc.GetCarts(&data); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *CartHandler) GetCartByIdHandler(c *gin.Context) {
	ID := c.Param("id")
	var data Cart

	if err := h.svc.GetCart(&data, ID); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *CartHandler) PostCartHandler(c *gin.Context) {
	var data Cart
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.CreateCart(&data); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
