package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	server "github.com/rohansingh9001/Neura-Launch-Dashboard/controllers"
	"github.com/rohansingh9001/Neura-Launch-Dashboard/helpers"
)

func main() {

	// Loading the env variables
	err := godotenv.Load("../.env")
	helpers.CheckError(err)

	// Initializig the gin router/engine
	router := server.Server{}
	router.Server()

	// Making an API route 
	router.PingRoute()

	// Starting the Server
	fmt.Println("Server started at port 8080")
	router.RunServer()

}
