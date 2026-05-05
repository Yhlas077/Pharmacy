package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /users
func UserCreate(c *gin.Context) {
	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.UserErrorResponse{err.Error(), "400"})
		return
	}

	_, err := utils.GetDB().Exec(c.Request.Context(),
		"INSERT INTO users(id, name, email, password, role) VALUES ($1,$2,$3,$4,$5)",
		req.ID, req.Name, req.Email, req.Password, req.Role,
	)

	if err != nil {
		c.JSON(500, models.UserErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// GET /users
func UserList(c *gin.Context) {
	rows, err := utils.GetDB().Query(c.Request.Context(),
		"SELECT id, name, email, password, role FROM users")

	if err != nil {
		c.JSON(500, models.UserErrorResponse{err.Error(), "500"})
		return
	}
	defer rows.Close()

	var list []models.User

	for rows.Next() {
		var e models.User

		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Role); err != nil {
			c.JSON(500, models.UserErrorResponse{err.Error(), "500"})
			return
		}

		list = append(list, e)
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /users/:id
func UserDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	res, err := utils.GetDB().Exec(c.Request.Context(),
		"DELETE FROM users WHERE id=$1", id)

	if err != nil || res.RowsAffected() == 0 {
		c.JSON(500, false)
		return
	}

	c.JSON(200, true)
}

// PUT /users/:id
func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, false)
		return
	}

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.UserErrorResponse{err.Error(), "400"})
		return
	}

	_, err = utils.GetDB().Exec(c.Request.Context(),
		"UPDATE users SET id=$1, name=$2, email=$3, password=$4, role=$5 WHERE id=$6",
		req.ID, req.Name, req.Email, req.Password, req.Role, id,
	)

	if err != nil {
		c.JSON(500, models.UserErrorResponse{err.Error(), "500"})
		return
	}

	c.JSON(200, true)
}

// ENDPOINT
func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", UserCreate)
	r.GET("/users", UserList)
	r.DELETE("/users/:id", UserDelete)
	r.PUT("/users/:id", UserUpdate)
}
