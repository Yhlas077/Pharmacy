package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"

	"github.com/gin-gonic/gin"
)

// POST /Category  // controllers
func CategoryCreate(c *gin.Context) {

	var req models.Categories

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.CategoriesErrorResponse{err.Error(), "400"})
		return
	}

	_, err := repositories.CategoryCreate(c.Request.Context(), req)

	if err != nil {
		c.JSON(500, models.CategoriesErrorResponse{err.Error(), "400"})
	}

	c.JSON(200, true)
}

// GET /Category
func CategoryList(c *gin.Context) {

	var filter repositories.CategoryFilter
	var list []models.Categories

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	list, err := repositories.CategoryList(c.Request.Context(), filter)

	if err != nil {
		c.JSON(400, false)
		return
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /Category/:id
func CategoryDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.CategoryDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

// PUT /Category/:id
func CategoryUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var req models.Categories

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.CategoryUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}

// ENDPOINT
func CategoryRoutes(r *gin.Engine) {
	r.POST("/category", CategoryCreate)
	r.GET("/category", CategoryList)
	r.DELETE("/category/:id", CategoryDelete)
	r.PUT("/category/:id", CategoryUpdate)
}
