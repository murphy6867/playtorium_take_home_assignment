package cart_item

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CartItemHandler struct {
	svc CartItemService
}

func NewCartItemHandler(svc CartItemService) *CartItemHandler {
	return &CartItemHandler{svc: svc}
}

func (h *CartItemHandler) CreateCartItemsHandler(c *gin.Context) {
	var data CartItem
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Body is not valid",
		})
		return
	}

	if err := h.svc.CreateCartItemService(&data); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, data)
}

func (h *CartItemHandler) GetCartItemByCartIDHandler(c *gin.Context) {
	cartID := c.Param("cartID")
	data := make([]CartItem, 0)

	if err := h.svc.GetCartItemByCartIDService(&data, cartID); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
