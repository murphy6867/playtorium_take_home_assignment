package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
)

type DomainError struct {
	Code    int
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

func NewDomainError(code int, message string) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
	}
}

func HandleError(c *gin.Context, err error) {
	var domainErr *DomainError
	if errors.As(err, &domainErr) {
		c.JSON(domainErr.Code, gin.H{"error": domainErr.Message})
		return
	}

	log.Printf("Unhandled error: %v", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
}

func RoundFloat(val float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(val*pow) / pow
}
