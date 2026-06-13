package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/models"
)

type ErrorCode string

const ErrorCodeRequired ErrorCode = "required"
const ErrorCodeInvalid ErrorCode = "invalid"
const ErrorCodeMax ErrorCode = "max"
const ErrorCodeForbidden ErrorCode = "forbidden"

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
		"data":    data,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Query("token")
		userID := TokenMap[token]

		if c.Request.URL.Path != "/api/login" &&
			c.Request.URL.Path != "/api/registration" &&
			c.Request.URL.Path != "/api/logout" {
			if userID == 0 {
				ErrorResponse(c, errors.New("token is missing"), 400, ErrorCodeRequired)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
