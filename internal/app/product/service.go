package product

import (
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func ServiceGetProducts(c *gin.Context, data *[]Product) error {
	if err := RepositoryGetProducts(c, data); err != nil {
		return err
	}

	return nil
}

func ServiceGetProduct(c *gin.Context, data *Product, id string) error {
	if id == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Product id is required")
	}

	if err := RepositoryGetProduct(c, data, id); err != nil {
		return err
	}

	return nil
}

func ServiceCreateProduct(c *gin.Context, data *Product) error {
	if data.Name == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Category name is required")
	}

	if data.Price == 0 || data.Price < 0 {
		return utils.NewDomainError(http.StatusBadRequest, "Price is required")
	}

	if data.CategoryName == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Category name is required")
	}

	if err := RepositoryCreatProduct(c, data); err != nil {
		return err
	}

	return nil
}
