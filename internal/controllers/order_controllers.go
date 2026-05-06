package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"

	"github.com/gin-gonic/gin"
)

// POST /Orders  // controllers
func OrdersCreate(c *gin.Context) {

	var req models.Orders

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.OrdersErrorResponse{err.Error(), "400"})
		return
	}

	_, err := repositories.OrdersCreate(c.Request.Context(), req)

	if err != nil {
		c.JSON(500, models.OrdersErrorResponse{err.Error(), "400"})
	}

	c.JSON(200, true)
}

// GET /Orders
func OrdersList(c *gin.Context) {

	var filter repositories.OrdersFilter
	var list []models.Orders

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	list, err := repositories.OrdersList(c.Request.Context(), filter)

	if err != nil {
		c.JSON(400, false)
		return
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /Orders/:id
func OrdersDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.OrdersDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

// PUT /Orders/:id
func OrdersUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var req models.Orders

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.OrdersUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}

// ENDPOINT
func OrdersRoutes(r *gin.Engine) {
	r.POST("/orders", OrdersCreate)
	r.GET("/orders", OrdersList)
	r.DELETE("/orders/:id", OrdersDelete)
	r.PUT("/orders/:id", OrdersUpdate)
}
