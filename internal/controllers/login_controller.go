package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

var TokenMap map[string]int
var email string
var password string

func Login(c *gin.Context) {

	email = c.Query("email")
	password = c.Query("password")

	info, err := repositories.UserGetByEmail(c, email)

	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
	}

	if info.Password == password {
		info, _ := repositories.UserGetByEmail(context.Background(), email)
		if TokenMap == nil {
			TokenMap = map[string]int{}
		}
		token := GenerateToken(email)
		TokenMap[token] = info.ID
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		utils.ErrorResponse(c, err, 500, "")
	}
}

func LoginRoute(r *gin.Engine) {
	r.POST("/login", Login)
}
