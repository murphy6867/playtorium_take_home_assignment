package category

import (
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/config"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func RepositoryGetCategories(c *gin.Context, data *[]Category) error {
	if err := config.DB.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No category found")
	}

	return nil
}

func RepositoryCreatCategory(c *gin.Context, data *Category) error {
	if err := config.DB.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}
