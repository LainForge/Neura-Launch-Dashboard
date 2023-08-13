package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rohansingh9001/Neura-Launch-Dashboard/helpers"
)

func main() {

	// Loading the env variables
	err := godotenv.Load("../.env")
	helpers.CheckError(err)

	// Initializig the gin router/engine
	router := gin.Default()

	// Making an API route 
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Starting the Server
	fmt.Println("Server started at port 8080")
	log.Fatal(router.Run(":8080"))

}
