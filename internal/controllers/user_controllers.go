package controllers

import (
	"strconv"
	"strings"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/services"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /users  // controllers
func UserCreate(c *gin.Context) {

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.UserCreate(c.Request.Context(), req)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
	}

	utils.SuccessResponse(c, nil, models.Meta{})
}

// DELETE /users/:id
func UserDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.UserDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil, models.Meta{})
}

// GET /users
func UserList(c *gin.Context) {

	var filter repositories.UserFilter
	var list []models.User

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	filter.Search = c.Query("search")
	filter.Role = c.Query("role")

	list, err := repositories.UserList(c.Request.Context(), filter)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var totalUsers int
	query := "SELECT COUNT(*) FROM users"
	err = repositories.GetDB().QueryRow(c, query).Scan(&totalUsers)

	utils.SuccessResponse(c, list, models.Meta{
		Total:  totalUsers,
		Limit:  filter.Limit,
		Offset: filter.Offset,
	})
}

// PUT /users/:id
func UserUpdate(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)

	var req models.UserUpdateRequest
	err := c.Bind(&req)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = services.UpdateUserService(c.Request.Context(), token, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
	utils.SuccessResponse(c, nil, models.Meta{})
}

func GetUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)
	req, err := services.GetUserService(c, token, false)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req, models.Meta{})
}

// ENDPOINT
func UserRoutes(rg *gin.RouterGroup) {
	rg.Group("").Use(utils.RequireUser())
	rg.POST("/admin/users", UserCreate)
	rg.GET("/admin/users", UserList)
	rg.DELETE("/admin/users/:id", UserDelete)
	rg.PUT("/admin/users/:id", UserUpdate)
	rg.GET("/admin/users/:id", GetUser)
}
