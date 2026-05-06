package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"

	"github.com/gin-gonic/gin"
)

// POST /users  // controllers
func UserCreate(c *gin.Context) {

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.UserErrorResponse{err.Error(), "400"})
		return
	}

	_, err := repositories.UserCreate(c.Request.Context(), req)

	if err != nil {
		c.JSON(500, models.UserErrorResponse{err.Error(), "400"})
	}

	c.JSON(200, true)
}

// GET /users
func UserList(c *gin.Context) {

	var filter repositories.UserFilter
	var list []models.User

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	filter.Search = c.Query("search")
	filter.Role = c.Query("role")

	list, err := repositories.UserList(c.Request.Context(), filter)

	if err != nil {
		c.JSON(400, false)
		return
	}

	c.JSON(200, gin.H{
		"list": list,
	})
}

// DELETE /users/:id
func UserDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.UserDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

// PUT /users/:id
func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.UserUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}

// ENDPOINT
func UserRoutes(r *gin.Engine) {
	r.POST("/users", UserCreate)
	r.GET("/users", UserList)
	r.DELETE("/users/:id", UserDelete)
	r.PUT("/users/:id", UserUpdate)
}
