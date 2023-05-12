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
	hostname := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", hostname, port)

	r := gin.Default()

	models.ConnectDB()

	public := r.Group("/api")
	public.GET("/users", controllers.GetUsers)
	public.GET("/users/:id", controllers.GetUser)
	public.PUT("/users/:id", controllers.UpdateUser)
	public.DELETE("/users/:id", controllers.DeleteUser)
	public.POST("/users", controllers.CreateUser)
	r.Run(addr)
}
