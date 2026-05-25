package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/controllers"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Query("token")

		userID := controllers.TokenMap[token]
		if c.Request.URL.Path != "/api/login" && c.Request.URL.Path != "/api/registration" && c.Request.URL.Path != "/api/logout" {
			if userID == 0 {
				utils.ErrorResponse(c, errors.New("token is missing"), 400, utils.ErrorCodeRequired)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// MAIN
func main() {

	utils.ConnectDB("postgres://postgres:123456@localhost:5432/pharmacy_db")

	defer utils.GetDB().Close(context.Background())

	r := gin.Default()

	r.Use(Logger())

	rg := r.Group("/api")

	controllers.UserRoutes(rg)
	controllers.PharmacyMedicineRoutes(rg)
	controllers.PharmacyRoutes(rg)
	controllers.OrdersRoutes(rg)
	controllers.CategoryRoutes(rg)
	controllers.LoginRoute(rg)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

}
