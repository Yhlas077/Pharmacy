package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yhlas/basic-pharmacy/internal/models"
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

func Login(c *gin.Context) {

	var token string

	email := c.Query("email")
	password := c.Query("password")

	Info, err := repositories.UserGetByEmail(c, email)

	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	if Info.Password == password {

		if utils.TokenMap == nil {
			utils.TokenMap = map[string]int{}
		}
		token = GenerateToken(email)
		utils.TokenMap[token] = Info.ID

		repositories.InsertToken(c, Info.ID, token)
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
}

func Registration(c *gin.Context) {

	name := c.Query("name")
	email := c.Query("email")
	password := c.Query("password")

	validate := validator.New()

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     "user",
	}

	err := validate.Struct(newUser)

	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	repositories.InsertUser(c, name, email, password)
	utils.SuccessResponse(c, nil)
}

func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	repositories.DeleteToken(c, token)
	utils.SuccessResponse(c, nil)
}

func LoginRoute(rg *gin.RouterGroup) {
	rg.POST("/login", Login)
	rg.POST("/registration", Registration)
	rg.DELETE("/logout", Logout)
}
