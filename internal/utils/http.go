package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/models"
)

type ErrorCode string

const ErrorCodeRequired ErrorCode = "required"
const ErrorCodeInvalid ErrorCode = "invalid"
const ErrorCodeMax ErrorCode = "max"

// Error
func ErrorResponse(c *gin.Context, err error, status int, errorCode ErrorCode) {
	c.JSON(status, models.ErrorResponse{
		Success:   false,
		ErrorMsg:  err.Error(),
		ErrorCode: string(errorCode),
	})
}

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"success": true,
		"error":   false,
		"data": data,
	})
}
