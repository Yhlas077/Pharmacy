package controllers

import (
	"strconv"

	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/services"
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

	err := services.CreateOrderService(c.Request.Context(), req.Name, req.Price, req.Description)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
	}
	utils.SuccessResponse(c, nil,models.Meta{})
}

// GET /Orders
func OrdersList(c *gin.Context) {

	var filter repositories.OrdersFilter
	var list []models.Orders

	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	filter.Search = c.Query("search")

	list, err := repositories.OrdersList(c.Request.Context(), filter)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var totalUsers int
	query := "SELECT COUNT(*) FROM orders"
	err = repositories.GetDB().QueryRow(c, query).Scan(&totalUsers)

	utils.SuccessResponse(c, list, models.Meta{
		Total: totalUsers,
		Limit: filter.Limit,
		Offset:filter.Offset,
	})

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
	utils.SuccessResponse(c, nil, models.Meta{})
}

// PUT /Orders/:id
func OrdersUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.OrderCreateRequest

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.OrdersUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, "")
		return
	}
	utils.SuccessResponse(c, nil, models.Meta{})
}

func GetOrder(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	req, err := services.GetOrderServices(c, id)
	if utils.ErrorCheck(c, err) {
		return
	}
	utils.SuccessResponse(c, req, models.Meta{})
}

// ENDPOINT
func OrdersRoutes(rg *gin.RouterGroup) {
	rg.POST("/admin/orders", OrdersCreate)
	rg.GET("/admin/orders", OrdersList)
	rg.DELETE("/admin/orders/:id", OrdersDelete)
	rg.PUT("/admin/orders/:id", OrdersUpdate)
	rg.GET("/admin/orders/get/:id", GetOrder)

}
