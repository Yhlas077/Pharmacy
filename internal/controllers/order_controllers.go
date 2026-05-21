package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /Orders  // controllers
func OrdersCreate(c *gin.Context) {

	var req models.Orders

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.OrdersCreate(c.Request.Context(), req)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
	}
	utils.SuccessResponse(c, nil)
}

// GET /Orders
func OrdersList(c *gin.Context) {

	var filter repositories.OrdersFilter
	var list []models.Orders

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	list, err := repositories.OrdersList(c.Request.Context(), filter)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, list)

}

// DELETE /Orders/:id
func OrdersDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.OrdersDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
	utils.SuccessResponse(c, nil)
}

// PUT /Orders/:id
func OrdersUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.Orders

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.OrdersUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
	utils.SuccessResponse(c, nil)
}

// ENDPOINT
func OrdersRoutes(rg *gin.RouterGroup) {
	rg.POST("/orders", OrdersCreate)
	rg.GET("/orders", OrdersList)
	rg.DELETE("/orders/:id", OrdersDelete)
	rg.PUT("/orders/:id", OrdersUpdate)
}
