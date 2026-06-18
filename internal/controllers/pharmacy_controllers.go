package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/services"
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

	err := services.CreatePharmacyService(c.Request.Context(), req.Name, req.Address, req.PharmacyHours)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}
	utils.SuccessResponse(c, nil, models.Meta{})
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

	var totalUsers int
	query := "SELECT COUNT(*) FROM pharmacies"
	err = repositories.GetDB().QueryRow(c, query).Scan(&totalUsers)

	utils.SuccessResponse(c, list, models.Meta{
		Total: totalUsers,
		Limit: filter.Limit,
		Offset:filter.Offset,
	})
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
	utils.SuccessResponse(c, nil, models.Meta{})
}

// PUT /Pharmacies/:id
func PharmacyUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.PharmacyCreateRequest

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.PharmacyUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil, models.Meta{})
}

func GetPharmacy(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := services.GetPharmacyService(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req, models.Meta{})
}

// ENDPOINT
func PharmacyRoutes(rg *gin.RouterGroup) {
	rg.POST("/admin/pharmacies", PharmacyCreate)
	rg.GET("/admin/pharmacies", PharmacyList)
	rg.DELETE("/admin/pharmacies/:id", PharmacyDelete)
	rg.PUT("/admin/pharmacies/:id", PharmacyUpdate)
	rg.GET("/admin/pharmacies/get/:id", GetPharmacy)
}
