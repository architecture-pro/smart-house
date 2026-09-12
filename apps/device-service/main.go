package main

import (
	"log"
	"os"
	"device-service/controllers"
	"device-service/db"
	"device-service/services"
	"github.com/gin-gonic/gin"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@localhost:5434/devices"
	}

	database, err := db.New(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	deviceService := services.NewDeviceService(database)

	registrationController := controllers.NewRegistrationController(deviceService)
	connectionController := controllers.NewConnectionController(deviceService)
	deviceController := controllers.NewDeviceController(deviceService)

	router := gin.Default()

	api := router.Group("/api/v1")

	registrationController.RegisterRoutes(api)
	connectionController.RegisterRoutes(api)
	deviceController.RegisterRoutes(api)

	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}