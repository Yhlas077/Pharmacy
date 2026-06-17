package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yhlas/basic-pharmacy/config"
	"github.com/yhlas/basic-pharmacy/internal/controllers"
	"github.com/yhlas/basic-pharmacy/internal/repositories"
	"github.com/yhlas/basic-pharmacy/internal/utils"
)

// MAIN
func main() {

	config.LoadConfig()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretConnectionText := os.Getenv("DB_URL")

	repositories.ConnectDB(secretConnectionText)

	defer repositories.GetDB().Close(context.Background())

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(utils.AuthMiddleware())

	rg := r.Group("/api")

	controllers.UserRoutes(rg)
	controllers.PharmacyMedicineRoutes(rg)
	controllers.PharmacyRoutes(rg)
	controllers.OrdersRoutes(rg)
	controllers.CategoryRoutes(rg)

	controllers.LoginRoute(rg)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

}
