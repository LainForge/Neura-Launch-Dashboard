package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/rohansingh9001/Neura-Launch-Dashboard/dashboard/controllers"
	"github.com/rohansingh9001/Neura-Launch-Dashboard/dashboard/initializers"
	"github.com/rohansingh9001/Neura-Launch-Dashboard/dashboard/middlewares"
)

func init() {
	initializers.LoanEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	// Initializig the gin router/engine
	r := gin.Default()

	// Allowing Cross Origin Requests
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Auth URLs
	r.POST("/signup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	// Generate Token
	r.POST("/token", middlewares.RequireAuth, controllers.TokenController)

	// Ping endpoint
	r.GET("/ping", controllers.Ping)
	r.GET("", controllers.Hello)

	// Starting the Server
	fmt.Println("Server started at port 8080")
	r.Run()

}
