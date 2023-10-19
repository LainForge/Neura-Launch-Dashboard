package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingController(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Builder Service",
		})
}
