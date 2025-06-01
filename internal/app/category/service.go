package category

import (
	"github.com/gin-gonic/gin"
	"github.com/murphy6867/productcheckout/internal/utils"
	"net/http"
)

func ServiceGetCategory(c *gin.Context, data *[]Category) error {
	if err := RepositoryGetCategories(c, data); err != nil {
		return err
	}

	return nil
}

func ServiceCreateCategory(c *gin.Context, data *Category) error {
	if data.Name == "" {
		return utils.NewDomainError(http.StatusBadRequest, "Category name is required")
	}

	if err := RepositoryCreatCategory(c, data); err != nil {
		return err
	}

	return nil
}
