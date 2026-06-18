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
func PharmacyMedicineCreate(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	token := strings.TrimPrefix(auth, "Bearer ")
	token = strings.TrimSpace(token)

	var req models.PharmacyMedicinesCreateRequest

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err := services.CreatePharmacyMedicineService(c.Request.Context(), req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId, token)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
	}

	utils.SuccessResponse(c, nil)
}

// GET /users
func PharmacyMedicineList(c *gin.Context) {

	var filter repositories.PharmacyMedicineFilter
	var list []models.PharmacyMedicines

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	filter.Search = c.Query("search")

	list, err := repositories.PharmacyMedicineList(c.Request.Context(), filter)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, list)

}

// DELETE /users/:id
func PharmacyMedicineDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.PharmacyMedicineDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil)
}

// PUT /users/:id
func PharmacyMedicineUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.PharmacyMedicinesCreateRequest

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.PharmacyMedicineUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}

	utils.SuccessResponse(c, nil)
}

func GetPharmacyMedicine(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := services.GetPharmacyMedicineServices(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req)
}

// ENDPOINT
func PharmacyMedicineRoutes(rg *gin.RouterGroup) {
	rg.Group("").Use(utils.RequirePharmacyAdmin())
	rg.POST("/admin/pharmacy_medicine", PharmacyMedicineCreate)
	rg.GET("/admin/pharmacy_medicine", PharmacyMedicineList)
	rg.DELETE("/admin/pharmacy_medicine/:id", PharmacyMedicineDelete)
	rg.PUT("/admin/pharmacy_medicine/:id", PharmacyMedicineUpdate)
	rg.GET("/admin/pharmacy_medicines/get/:id", GetPharmacyMedicine)

}
