package repositories

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

func ShowToken(c *gin.Context, token string) {
	c.JSON(200, gin.H{
		"token": token,
	})
}

func InsertToken(c *gin.Context, id int, token string) {
	_, err := utils.GetDB().Exec(context.Background(), "INSERT into tokens(user_id, token) values ($1, $2)", id, token)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}
}

func GetToken(c *gin.Context, token string, id int) {
	row := utils.GetDB().QueryRow(context.Background(), `select token from tokens where id =$1`, id)
	row.Scan(&token)
}

func DeleteToken(c *gin.Context, token string) {
	var id int
	row := utils.GetDB().QueryRow(context.Background(), `select user_id from tokens where token = $1`, token)
	row.Scan(&id)
	_, err := utils.GetDB().Exec(c, `DELETE FROM tokens WHERE user_id=$1`, id)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeInvalid)
		return
	}
}

func InsertUser(c *gin.Context, name string, email string, password string) {
	_, err := utils.GetDB().Exec(context.Background(),
		"INSERT INTO users(name, email, password, role) VALUES ($1,$2,$3,$4)",
		name, email, password, "user",
	)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}
}
