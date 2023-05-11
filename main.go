package main

import (
	"fmt"
	"log"
	"os"

	"github.com/b4sile/bike-service-backend/internal/controllers"
	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", hostname, port)

	r := gin.Default()

	models.ConnectDB()

	r.GET("/ping", controllers.Ping)
	r.Run(addr)
}
