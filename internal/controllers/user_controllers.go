package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
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

	utils.SuccessResponse(c, nil)
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

	utils.SuccessResponse(c, list)
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

	utils.SuccessResponse(c, nil)
}

// PUT /users/:id
func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.UserUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
	utils.SuccessResponse(c, nil)
}

// ENDPOINT
func UserRoutes(rg *gin.RouterGroup) {
	rg.Group("").Use(utils.RequireUser())
	rg.POST("/users", UserCreate)
	rg.GET("/users", UserList)
	rg.DELETE("/users/:id", UserDelete)
	rg.PUT("/users/:id", UserUpdate)
}
