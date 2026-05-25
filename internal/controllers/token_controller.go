package controllers

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)

	return hex.EncodeToString(hasher.Sum(nil))
}

var TokenMap map[string]int

func Login(c *gin.Context) {

	var token string

	email := c.Query("email")
	password := c.Query("password")

	fmt.Println("Hello baby")

	Info, err := repositories.UserGetByEmail(c, email)

	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
	}

	if Info.Password == password {
		Info, _ := repositories.UserGetByEmail(context.Background(), email)
		if TokenMap == nil {
			TokenMap = map[string]int{}
		}
		token = GenerateToken(email)
		TokenMap[token] = Info.ID

		repositories.InsertToken(c, Info.ID, token)

		repositories.GetToken(c, token)

	} else {
		utils.ErrorResponse(c, err, 500, "")
	}
}

func LoginRoute(rg *gin.RouterGroup) {
	rg.POST("/login", Login)
}
