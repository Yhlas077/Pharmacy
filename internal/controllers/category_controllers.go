package controllers

import (
	"errors"
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/services"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	categoryidstr := c.Param("id")
	categoryid, _ := strconv.Atoi(categoryidstr)
	req, err := services.GetCategoryService(c, categoryid)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req, models.Meta{})
}

// POST /Category  // controllers
func CategoryCreate(c *gin.Context) {

	var req models.CategoryCreateRequest

	err := c.BindJSON(&req)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = services.CreateCategoryService(c, req.Name)

	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil, models.Meta{})
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

	var totalUsers int
	query := "SELECT COUNT(*) FROM categories"
	err = repositories.GetDB().QueryRow(c, query).Scan(&totalUsers)

	utils.SuccessResponse(c, list, models.Meta{
		Total: totalUsers,
		Limit: filter.Limit,
		Offset:filter.Offset,
	})
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

	utils.SuccessResponse(c, nil, models.Meta{})
}

// PUT /Category/:id
func CategoryUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.CategoryCreateRequest

	err = c.BindJSON(&req)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = services.UpdateCategoryService(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, nil, models.Meta{})
}

// ENDPOINT
func CategoryRoutes(rg *gin.RouterGroup) {
	rg.Group("").Use(utils.RequireAdmin())
	rg.POST("/admin/categories", CategoryCreate)
	rg.GET("/admin/categories", CategoryList)
	rg.DELETE("/admin/categories/:id", CategoryDelete)
	rg.PUT("/admin/categories/:id", CategoryUpdate)
	rg.GET("/admin/categories/get/:id", GetCategory)
}
