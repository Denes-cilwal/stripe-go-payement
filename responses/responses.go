package responses

import (
	"github.com/gin-gonic/gin"
	"stripe-go-payment/errors"
)

func HandleError(c *gin.Context, err error) {
	errorType := errors.GetErrorType(err)
	status := errors.GetStatusCode(errorType)

	errorContext := errors.GetErrorContext(err)
	if errorContext != nil {
		c.JSON(status, gin.H{
			"error":   err.Error(),
			"context": errorContext,
		})
		return
	}

	c.JSON(status, gin.H{"error": err.Error()})
}
