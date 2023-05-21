package main

import (
	"fmt"
	"log"
	"os"

	"github.com/b4sile/bike-service-backend/internal/controllers"
	"github.com/b4sile/bike-service-backend/internal/middlewares"
	"github.com/b4sile/bike-service-backend/internal/models"
	"github.com/gin-gonic/contrib/static"
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
	models.CreateAdminUser()

	public := r.Group("/api")
	auth := r.Group("/api").Use(middlewares.JwtAuthMiddleware())
	admin := r.Group("/api").Use(middlewares.JwtAuthMiddleware()).Use(middlewares.AdminAuthMiddleware())

	r.Use(static.Serve("/", static.LocalFile("./assets/build", true)))

	admin.GET("/users", controllers.GetUsers)
	admin.GET("/users/:id", controllers.GetUser)
	auth.PUT("/users/:id", controllers.UpdateUser)
	admin.DELETE("/users/:id", controllers.DeleteUser)
	public.POST("/users", controllers.CreateUser)
	auth.GET("/users/getMe", controllers.GetCurrentUser)

	public.GET("/services", controllers.GetServices)
	public.GET("/services/:id", controllers.GetService)
	admin.PUT("/services/:id", controllers.UpdateService)
	admin.DELETE("/services/:id", controllers.DeleteService)
	admin.POST("/services", controllers.CreateService)

	auth.GET("/requests", controllers.GetRequests)
	admin.GET("/requests/:id", controllers.GetRequest)
	auth.PUT("/requests/:id", controllers.UpdateRequest)
	auth.DELETE("/requests/:id", controllers.DeleteRequest)
	auth.POST("/requests", controllers.CreateRequest)

	auth.GET("/notifications", controllers.GetNotifications)
	admin.GET("/notifications/:id", controllers.GetNotification)
	admin.PUT("/notifications/:id", controllers.UpdateNotification)
	admin.DELETE("/notifications/:id", controllers.DeleteNotification)
	admin.POST("/notifications", controllers.CreateNotification)

	public.POST("/login", controllers.Login)

	r.Run(addr)
}
