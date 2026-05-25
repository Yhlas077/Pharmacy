package repositories

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

func GetToken(c *gin.Context, token string) {
	c.JSON(200, gin.H{
		"token": token,
	})
}

func InsertToken(c *gin.Context, id int, token string) {
	_, err := utils.GetDB().Exec(context.Background(), "INSERT into token(userid, token) values ($1, $2)", id, token)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}
}
