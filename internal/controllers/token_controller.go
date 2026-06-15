package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"

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

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	email := req.Email
	password := req.Password

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

		utils.SuccessResponse(c, gin.H{
			"token":      token,
			"expires_at": time.Now().AddDate(1, 0, 0),
			"user_id":    Info.ID,
		})

		repositories.InsertToken(c, Info.ID, token)
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
}

func ChangePassword(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	var passchange models.ChangePasswordRequest
	err := c.Bind(&passchange)
	if utils.ErrorCheck(c, err) {
		return
	}
	var req models.User
	err = utils.ChangePassword(c, token, true, passchange, req)
	if utils.ErrorCheck(c, err) {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
	utils.SuccessResponse(c, "password changed")
}

func Registration(c *gin.Context) {

	type RegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	name := req.Name
	email := req.Email
	password := req.Password

	validate := validator.New()

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     "user",
	}

	err := validate.Struct(newUser)

	if err != nil {
		fmt.Println("VALIDATION ERROR:", err)
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
