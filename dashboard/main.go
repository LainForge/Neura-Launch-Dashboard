package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rohansingh9001/Neura-Launch-Dashboard/controllers"
	"github.com/rohansingh9001/Neura-Launch-Dashboard/initializers"
	"github.com/rohansingh9001/Neura-Launch-Dashboard/middlewares"
)

func init() {
	initializers.LoanEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {


	// Initializig the gin router/engine
	r := gin.Default()

	// Auth URLs
	r.POST("/signup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	// Ping endpoint
	r.GET("/ping", controllers.Ping)
	r.GET("", controllers.Hello)

	// Starting the Server
	fmt.Println("Server started at port 8080")
	r.Run()

}
