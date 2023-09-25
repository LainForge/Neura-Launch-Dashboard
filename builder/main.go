package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Routes
	r.GET("/verifyCheckSum", nil)

	// Starting the server
	r.Run(":8080")
}
