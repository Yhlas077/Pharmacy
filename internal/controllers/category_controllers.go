package controllers

import (
	"errors"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// IMPLEMENT: service folder for all endpoints
// example: CategoryCreate(req models.CategoryRequest) res models.CategoryResponse
// example in controller call (insted of repository): service.CategoryCreate(req)

// POST /Category  // controllers
func CategoryCreate(c *gin.Context) {

	var req models.Category

	err := c.BindJSON(&req)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err = repositories.CategoryCreate(c.Request.Context(), req)

	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil)
}

// GET /Category
func CategoryList(c *gin.Context) {

	var filter repositories.CategoryFilter
	var list []models.Category

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	if filter.Limit == 0 {
		utils.ErrorResponse(c, errors.New("limit"), 400, utils.ErrorCodeRequired)
		return
	}

	list, err := repositories.CategoryList(c.Request.Context(), filter)

	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, list)
}

// DELETE /Category/:id
func CategoryDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.CategoryDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil)
}

// PUT /Category/:id
func CategoryUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.Category

	err = c.BindJSON(&req)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.CategoryUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, nil)
}

// ENDPOINT
func CategoryRoutes(rg *gin.RouterGroup) {
	rg.Group("").Use(utils.RequireAdmin())
	rg.POST("/admin/category", CategoryCreate)
	rg.GET("/admin/category", CategoryList)
	rg.DELETE("/admin/category/:id", CategoryDelete)
	rg.PUT("/admin/category/:id", CategoryUpdate)
}
