package server

import "github.com/gin-gonic/gin"

// Server struct 
type Server struct {
	Router *gin.Engine
}

// Initialize the server
func (server *Server) Server() {
	server.Router = gin.Default()
}

func (server *Server) RunServer() {
	server.Router.Run(":8080")
}