package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yhlas/basic-pharmacy/internal/controllers"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

// MAIN
func main() {

	utils.ConnectDB("postgres://yhlas1:123456@localhost:5432/postgres")
	defer utils.GetDB().Close(context.Background())

	// HTTP serve
	r := gin.Default()

	controllers.UserRoutes(r)
	controllers.PharmacyMedicineRoutes(r)
	controllers.PharmacyRoutes(r)
	controllers.OrdersRoutes(r)
	controllers.CategoryRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
