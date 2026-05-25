package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /Pharmacies  // controllers
func PharmacyCreate(c *gin.Context) {

	var req models.Pharmacies

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.PharmacyCreate(c.Request.Context(), req)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
	}
	utils.SuccessResponse(c, nil)
}

// GET /Pharmacies
func PharmacyList(c *gin.Context) {

	var filter repositories.PharmacyFilter
	var list []models.Pharmacies

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	filter.Search = c.Query("search")

	list, err := repositories.PharmacyList(c.Request.Context(), filter)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, list)
}

// DELETE /Pharmacies/:id
func PharmacyDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.PharmacyDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
	utils.SuccessResponse(c, nil)
}

// PUT /Pharmacies/:id
func PharmacyUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.Pharmacies

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.PharmacyUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil)
}

// ENDPOINT
func PharmacyRoutes(rg *gin.RouterGroup) {
	rg.POST("/admin/pharmacies", PharmacyCreate)
	rg.GET("/admin/pharmacies", PharmacyList)
	rg.DELETE("/admin/pharmacies/:id", PharmacyDelete)
	rg.PUT("/admin/pharmacies/:id", PharmacyUpdate)
}
