package controllers

import (
	"github.com/yhlas/basic-pharmacy/internal/models"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"

	"github.com/gin-gonic/gin"
)

// POST /users  // controllers
func Registration(c *gin.Context) {

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.Registration(c.Request.Context(), req)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
	}

	utils.SuccessResponse(c, nil)
}

// ENDPOINT
func Post(rg *gin.RouterGroup) {
	rg.POST("/registration", Registration)
}
