package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /orders
func OrderCreate(c *gin.Context) {
	var req models.Orders

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.OrdersErrorResponse{err.Error(), "400"})
		return
	}

	_, err := utils.GetDB().Exec(c.Request.Context(),
		"INSERT INTO orders(id, name, price, description) VALUES ($1,$2,$3,$4)",
		req.ID, req.Name, req.Price, req.Description,
	)

	if err != nil {
		c.JSON(500, models.OrdersErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// GET /orders
func OrderList(c *gin.Context) {
	rows, err := utils.GetDB().Query(c.Request.Context(),
		"SELECT id, name, price, description FROM orders")

	if err != nil {
		c.JSON(500, models.OrdersErrorResponse{err.Error(), "500"})
		return
	}
	defer rows.Close()

	var list []models.Orders

	for rows.Next() {
		var e models.Orders

		if err := rows.Scan(&e.ID, &e.Name, &e.Price, &e.Description); err != nil {
			c.JSON(500, models.OrdersErrorResponse{err.Error(), "500"})
			return
		}

		list = append(list, e)
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /orders/:id
func OrderDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	res, err := utils.GetDB().Exec(c.Request.Context(),
		"DELETE FROM orders WHERE id=$1", id)

	if err != nil || res.RowsAffected() == 0 {
		c.JSON(500, false)
		return
	}

	c.JSON(200, true)
}

// PUT /orders/:id
func OrderUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	var req models.Orders

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.OrdersErrorResponse{err.Error(), "400"})
		return
	}

	_, err = utils.GetDB().Exec(c.Request.Context(),
		"UPDATE orders SET id=$1, name=$2, price=$3, description=$4, WHERE id=$5",
		req.ID, req.Name, req.Price, req.Description, id,
	)

	if err != nil {
		c.JSON(500, models.OrdersErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// ENDPOINT
func OrderRoutes(r *gin.Engine) {
	r.POST("/orders", OrderCreate)
	r.GET("/orders", OrderList)
	r.DELETE("/orders/:id", OrderDelete)
	r.PUT("/orders/:id", OrderUpdate)
}
