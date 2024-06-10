package main

import (
    "github.com/gin-gonic/gin"

    "hotel/pkg/config"
    "hotel/pkg/handlers"
	"hotel/middleware"
)

func main() {
    r := gin.Default()
    config.ConnectDB()
    r.POST("/users", handlers.CreateUser)

	r.POST("/login", handlers.Login)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/protected-resource", handlers.ProtectedResource)
	}

    r.Run()
}