package product

import (
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/config"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func RepositoryGetProducts(c *gin.Context, data *[]Product) error {
	if err := config.DB.Find(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No product found")
	}

	return nil
}

func RepositoryGetProduct(c *gin.Context, data *Product, id string) error {
	if err := config.DB.First(&data, id).Error; err != nil {
		return utils.NewDomainError(http.StatusNotFound, "No product found")
	}

	return nil
}

func RepositoryCreatProduct(c *gin.Context, data *Product) error {
	if err := config.DB.Create(&data).Error; err != nil {
		return utils.NewDomainError(http.StatusNotImplemented, "The request method is not supported by the server")
	}

	return nil
}
