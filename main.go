package main

import (
	"context"
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

		if c.Request.URL.Path != "/login" {
			if userID == 0 {
				c.JSON(400, gin.H{
					"error": "token missing",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// MAIN
func main() {

	utils.ConnectDB("postgres://yhlas1:123456@localhost:5432/postgres")
	defer utils.GetDB().Close(context.Background())

	// HTTP serve
	r := gin.Default()

	r.Use(Logger())

	controllers.UserRoutes(r)
	controllers.PharmacyMedicineRoutes(r)
	controllers.PharmacyRoutes(r)
	controllers.OrdersRoutes(r)
	controllers.CategoryRoutes(r)
	controllers.LoginRoute(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

}