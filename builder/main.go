package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohansingh9001/Neura-Launch-Dashboard/builder/controllers"
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/verifyCheckSum", controllers.VerifyCheckSumControllers)
	r.GET("/downloadFiles", nil)
	r.GET("/buildImage", nil)

	// Starting the server
	r.Run(":8080")
}
