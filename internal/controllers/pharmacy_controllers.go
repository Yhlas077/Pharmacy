package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /pharmacy
func PharmacyCreate(c *gin.Context) {
	var req models.Pharmacies

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.PharmacyErrorResponse{err.Error(), "400"})
		return
	}

	_, err := utils.GetDB().Exec(c.Request.Context(),
		"INSERT INTO pharmacies(id, name, address, pharmacy_hours) VALUES ($1,$2,$3,$4)",
		req.ID, req.Name, req.Address, req.Pharmacy_hours,
	)

	if err != nil {
		c.JSON(500, models.PharmacyErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// GET /pharmacy
func PharmacyList(c *gin.Context) {
	rows, err := utils.GetDB().Query(c.Request.Context(),
		"SELECT id, name, address, pharmacy_hours FROM pharmacies")

	if err != nil {
		c.JSON(500, models.PharmacyErrorResponse{err.Error(), "500"})
		return
	}
	defer rows.Close()

	var list []models.Pharmacies

	for rows.Next() {
		var e models.Pharmacies

		if err := rows.Scan(&e.ID, &e.Name, &e.Address, &e.Pharmacy_hours); err != nil {
			c.JSON(500, models.PharmacyErrorResponse{err.Error(), "500"})
			return
		}

		list = append(list, e)
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /pharmacies/:id
func PharmacyDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	res, err := utils.GetDB().Exec(c.Request.Context(),
		"DELETE FROM pharmacies WHERE id=$1", id)

	if err != nil || res.RowsAffected() == 0 {
		c.JSON(500, false)
		return
	}

	c.JSON(200, true)
}

// PUT /phramacies/:id
func PharmacyUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	var req models.Pharmacies

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.PharmacyErrorResponse{err.Error(), "400"})
		return
	}

	_, err = utils.GetDB().Exec(c.Request.Context(),
		"UPDATE pharmacies SET id=$1, name=$2, address=$3, pharmacy_hours=$4 WHERE id=$5",
		req.ID, req.Name, req.Address, req.Pharmacy_hours, id,
	)

	if err != nil {
		c.JSON(500, models.PharmacyErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// ENDPOINT
func PharmacyRoutes(r *gin.Engine) {
	r.POST("/pharmacy", PharmacyCreate)
	r.GET("/pharmacy", PharmacyList)
	r.DELETE("/pharmacy/:id", PharmacyDelete)
	r.PUT("/pharmacy/:id", PharmacyUpdate)
}