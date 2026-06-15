package utils

import (
	"context"
	"errors"

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

func ChangePassword(c context.Context, token string, word bool, passchange models.ChangePasswordRequest, req models.User) error {
	req, err := repositories.GetUser(c, token, word)
	if err != nil {
		return err
	}
	if req.Password != passchange.OldPassword {
		return errors.New("wrong pass")
	}
	err = UpdatePassword(c, token, passchange)
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
