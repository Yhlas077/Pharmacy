package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"

	"github.com/gin-gonic/gin"
)

// POST /Pharmacies  // controllers
func PharmacyCreate(c *gin.Context) {

	var req models.Pharmacies

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.PharmacyErrorResponse{err.Error(), "400"})
		return
	}

	_, err := repositories.PharmacyCreate(c.Request.Context(), req)

	if err != nil {
		c.JSON(500, models.PharmacyErrorResponse{err.Error(), "400"})
	}

	c.JSON(200, true)
}

// GET /Pharmacies
func PharmacyList(c *gin.Context) {

	var filter repositories.PharmacyFilter
	var list []models.Pharmacies

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	list, err := repositories.PharmacyList(c.Request.Context(), filter)

	if err != nil {
		c.JSON(400, false)
		return
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /Pharmacies/:id
func PharmacyDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.PharmacyDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

// PUT /Pharmacies/:id
func PharmacyUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var req models.Pharmacies

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.PharmacyUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}

// ENDPOINT
func PharmacyRoutes(r *gin.Engine) {
	r.POST("/pharmacies", PharmacyCreate)
	r.GET("/pharmacies", PharmacyList)
	r.DELETE("/pharmacies/:id", PharmacyDelete)
	r.PUT("/pharmacies/:id", PharmacyUpdate)
}
