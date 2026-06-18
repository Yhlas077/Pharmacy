package utils

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
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

func UpdatePassword(c context.Context, token string, passchange models.ChangePasswordRequest) error {
	db := repositories.GetDB()
	_, err := db.Exec(c, "update users u set password=$1 from tokens t where t.userid=u.id and t.token=$2", passchange.NewPassword, token)
	if err != nil {
		return err
	}
	return nil
}

func ErrorCheck(c *gin.Context, err error) bool {
	if err != nil {
		ErrorResponse(c, err, 401, ErrorCodeForbidden)
		return true
	}
	return false
}

func SuccessResponse(c *gin.Context, data any, meta models.Meta) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"meta":    models.Meta{
			Total: meta.Total,
			Limit: meta.Limit,
			Offset: meta.Offset,
		},
	})
}

func LoginSuccess(c *gin.Context, token string, expiry time.Time, Info models.User) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token":      token,
			"expires_at": expiry.Format(time.RFC3339),
			"user": gin.H{
				"id":    Info.ID,
				"name":  Info.Name,
				"email": Info.Email,
				"role":  Info.Role,
			},
		},
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		token := strings.TrimPrefix(auth, "Bearer ")
		token = strings.TrimSpace(token)

		if c.Request.URL.Path == "/api/auth/login" ||
			c.Request.URL.Path == "/api/register" ||
			c.Request.URL.Path == "/api/logout" {
			c.Next()
			return
		}

		var userID int
		query := "SELECT user_id FROM tokens WHERE token = $1"
		err := repositories.GetDB().QueryRow(c, query, token).Scan(&userID)

		if err != nil || userID == 0 {
			ErrorResponse(c, errors.New("token is missing or invalid"), 400, ErrorCodeRequired)
			c.Abort()
			return
		}

		c.Next()
	}
}
