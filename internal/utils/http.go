package utils

import (
	"context"
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
		"meta": models.Meta{
			Total:  meta.Total,
			Limit:  meta.Limit,
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
	publicPaths := map[string]bool{
		"/api/auth/login":    true,
		"/api/auth/register": true,
		"/api/auth/logout":   true,
	}

	return func(c *gin.Context) {
		if publicPaths[c.Request.URL.Path] {
			c.Next()
			return
		}

		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Success: false, ErrorMsg: "authorization token required",
				ErrorCode: string(ErrorCodeRequired),
			})
			c.Abort()
			return
		}
		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		if token == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Success: false, ErrorMsg: "authorization token required",
				ErrorCode: string(ErrorCodeRequired),
			})
			c.Abort()
			return
		}

		ctx := c.Request.Context()

		userID, err := repositories.GetUserIdByToken(ctx, token)
		if err != nil || userID == 0 {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Success: false, ErrorMsg: "invalid or missing token",
				ErrorCode: string(ErrorCodeForbidden),
			})
			c.Abort()
			return
		}

		expiresAt, err := repositories.GetExpiresAtByToken(ctx, token)
		if err != nil || expiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Success: false, ErrorMsg: "token expired",
				ErrorCode: string(ErrorCodeForbidden),
			})
			c.Abort()
			return
		}

		if TokenMap == nil {
			TokenMap = map[string]int{}
		}
		TokenMap[token] = userID

		c.Set("userID", userID)
		c.Set("token", token)

		c.Next()
	}
}
