package category

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

type CategoryHandler struct {
	svc CategoryService
}

func NewCategoryHandler(svc CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	category := make([]Category, 0)

	if err := h.svc.ServiceGetCategories(&category); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	var category Category

	if err := h.svc.ServiceGetCategory(&category, id); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) PostCategory(c *gin.Context) {
	var body Category

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
		})
		return
	}

	if err := h.svc.ServiceCreateCategory(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, body)
}
