package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /categories
func CategoryCreate(c *gin.Context) {
	var req models.Categories

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.CategoriesErrorResponse{err.Error(), "400"})
		return
	}

	_, err := utils.GetDB().Exec(c.Request.Context(),
		"INSERT INTO categories(id, name) VALUES ($1,$2)",
		req.ID, req.Name,
	)

	if err != nil {
		c.JSON(500, models.CategoriesErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// GET /categories
func CategoryList(c *gin.Context) {
	rows, err := utils.GetDB().Query(c.Request.Context(),
		"SELECT id, name from categories")

	if err != nil {
		c.JSON(500, models.CategoriesErrorResponse{err.Error(), "500"})
		return
	}
	defer rows.Close()

	var list []models.Categories

	for rows.Next() {
		var e models.Categories

		if err := rows.Scan(&e.ID, &e.Name); err != nil {
			c.JSON(500, models.CategoriesErrorResponse{err.Error(), "500"})
			return
		}

		list = append(list, e)
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /users/:id
func CategoryDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	res, err := utils.GetDB().Exec(c.Request.Context(),
		"DELETE FROM categories WHERE id=$1", id)

	if err != nil || res.RowsAffected() == 0 {
		c.JSON(500, false)
		return
	}

	c.JSON(200, true)
}

// PUT /users/:id
func CategoryUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	var req models.Categories

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.CategoriesErrorResponse{err.Error(), "400"})
		return
	}

	_, err = utils.GetDB().Exec(c.Request.Context(),
		"UPDATE categories SET id=$1, name=$2 WHERE id=$3",
		req.ID, req.Name, id,
	)

	if err != nil {
		c.JSON(500, models.CategoriesErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// ENDPOINT
func CategoryRoutes(r *gin.Engine) {
	r.POST("/categories", CategoryCreate)
	r.GET("/categories", CategoryList)
	r.DELETE("/categories/:id", CategoryDelete)
	r.PUT("/categories/:id", CategoryUpdate)
}
