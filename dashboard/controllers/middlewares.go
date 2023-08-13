package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func(server *Server) Ping(context *gin.Context) {
	context.JSON(
		http.StatusOK, 
		gin.H{
		"message": "Hello World!",
	})
}