package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
)

var TokenMap map[string]int

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ErrorResponse(c, errors.New("Authorization token required"), 401, ErrorCodeRequired)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userID := TokenMap[token]
		if userID == 0 {
			ErrorResponse(c, errors.New("token missing"), 403, ErrorCodeForbidden)
			c.Abort()
			return
		}
		user, err := repositories.UserGetByID(c, userID)
		if err != nil {
			ErrorResponse(c, err, 500, ErrorCodeForbidden)
			c.Abort()
			return
		}
		if user.Role != string(models.AdminRole) {
			ErrorResponse(c, errors.New("forbidden"), 403, ErrorCodeForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequirePharmacyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ErrorResponse(c, errors.New("Authorization token required"), 401, ErrorCodeRequired)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID := TokenMap[token]

		if userID == 0 {
			ErrorResponse(c, errors.New("token missing"), 403, ErrorCodeForbidden)
			c.Abort()
			return
		}
		user, err := repositories.UserGetByID(c, userID)
		if err != nil {
			ErrorResponse(c, err, 500, ErrorCodeForbidden)
			c.Abort()
			return
		}
		if user.Role != string(models.PharmacyRole) && user.Role != string(models.AdminRole) {
			ErrorResponse(c, errors.New("forbidden"), 403, ErrorCodeForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ErrorResponse(c, errors.New("Authorization token required"), 401, ErrorCodeRequired)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userID := TokenMap[token]
		if userID == 0 {
			ErrorResponse(c, errors.New("token missing"), 403, ErrorCodeForbidden)
			c.Abort()
			return
		}

		user, err := repositories.UserGetByID(c, userID)
		if err != nil {
			ErrorResponse(c, err, 500, ErrorCodeForbidden)
			c.Abort()
			return
		}

		c.Set("currentUser", user)

		c.Next()
	}
}
