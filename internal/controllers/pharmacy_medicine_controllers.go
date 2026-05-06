package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"

	"github.com/gin-gonic/gin"
)

// POST /users  // controllers
func PharmacyMedicineCreate(c *gin.Context) {

	var req models.PharmacyMedicines

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.PharmacyMedicineErrorResponse{err.Error(), "400"})
		return
	}

	_, err := repositories.PharmacyMedicineCreate(c.Request.Context(), req)

	if err != nil {
		c.JSON(500, models.PharmacyMedicineErrorResponse{err.Error(), "400"})
	}

	c.JSON(200, true)
}

// GET /users
func PharmacyMedicineList(c *gin.Context) {

	var filter repositories.PharmacyMedicineFilter
	var list []models.PharmacyMedicines

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	list, err := repositories.PharmacyMedicineList(c.Request.Context(), filter)

	if err != nil {
		c.JSON(400, false)
		return
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /users/:id
func PharmacyMedicineDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.PharmacyMedicineDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

// PUT /users/:id
func PharmacyMedicineUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var req models.PharmacyMedicines

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.PharmacyMedicineUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}

// ENDPOINT
func PharmacyMedicineRoutes(r *gin.Engine) {
	r.POST("/pharmacy_medicine", PharmacyMedicineCreate)
	r.GET("/pharmacy_medicine", PharmacyMedicineList)
	r.DELETE("/pharmacy_medicine/:id", PharmacyMedicineDelete)
	r.PUT("/pharmacy_medicine/:id", PharmacyMedicineUpdate)
}
