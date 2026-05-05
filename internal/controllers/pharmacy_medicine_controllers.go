package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /pharmacy_medicines
func Pharmacy_medicines_Create(c *gin.Context) {
	var req models.Pharmacy_medicines

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.PharmacyMedicineErrorResponse{err.Error(), "400"})
		return
	}

	_, err := utils.GetDB().Exec(c.Request.Context(),
		"INSERT INTO pharmacy_medicines(id, name, description, price, new_price, category_id) VALUES ($1,$2,$3,$4,$5,$6)",
		req.ID, req.Name, req.Description, req.Price, req.New_price, req.Category_id,
	)

	if err != nil {
		c.JSON(500, models.PharmacyMedicineErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// GET /pharmacy_medicines
func Pharmacy_medicines_List(c *gin.Context) {
	rows, err := utils.GetDB().Query(c.Request.Context(),
		"SELECT id, name, description, price, new_price, category_id FROM pharmacy_medicines")

	if err != nil {
		c.JSON(500, models.PharmacyMedicineErrorResponse{err.Error(), "500"})
		return
	}
	defer rows.Close()

	var list []models.Pharmacy_medicines

	for rows.Next() {
		var e models.Pharmacy_medicines

		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Price, &e.New_price, &e.Category_id); err != nil {
			c.JSON(500, models.PharmacyMedicineErrorResponse{err.Error(), "500"})
			return
		}

		list = append(list, e)
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /pharmacy_medicines/:id
func Pharmacy_medicines_Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	res, err := utils.GetDB().Exec(c.Request.Context(),
		"DELETE FROM pharmacy_medicines WHERE id=$1", id)

	if err != nil || res.RowsAffected() == 0 {
		c.JSON(500, false)
		return
	}

	c.JSON(200, true)
}

// PUT /pharmacy_medicines/:id
func Pharmacy_medicines_Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	var req models.Pharmacy_medicines

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.PharmacyMedicineErrorResponse{err.Error(), "400"})
		return
	}

	_, err = utils.GetDB().Exec(c.Request.Context(),
		"UPDATE pharmacy_medicines SET id=$1, name=$2, description=$3, price=$4, new_price=$5, category_id=$6 WHERE id=$7",
		req.ID, req.Name, req.Description, req.Price, req.New_price, req.Category_id, id,
	)

	if err != nil {
		c.JSON(500, models.PharmacyMedicineErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// ENDPOINT
func Pharmacy_medicines_Routes(r *gin.Engine) {
	r.POST("/pharmacy_medicines", Pharmacy_medicines_Create)
	r.GET("/pharmacy_medicines", Pharmacy_medicines_List)
	r.DELETE("/pharmacy_medicines/:id", Pharmacy_medicines_Delete)
	r.PUT("/pharmacy_medicines/:id", Pharmacy_medicines_Update)
}
