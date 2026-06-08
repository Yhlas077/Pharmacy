package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /users  // controllers
func PharmacyMedicineCreate(c *gin.Context) {

	var req models.PharmacyMedicines

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.PharmacyMedicineCreate(c.Request.Context(), req)

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

	var req models.PharmacyMedicines

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

// ENDPOINT
func PharmacyMedicineRoutes(rg *gin.RouterGroup) {
	rg.POST("/admin/pharmacy_medicine", PharmacyMedicineCreate).Use(utils.RequirePharmacyAdmin())
	rg.GET("/admin/pharmacy_medicine", PharmacyMedicineList)
	rg.DELETE("/admin/pharmacy_medicine/:id", PharmacyMedicineDelete)
	rg.PUT("/admin/pharmacy_medicine/:id", PharmacyMedicineUpdate)
}
