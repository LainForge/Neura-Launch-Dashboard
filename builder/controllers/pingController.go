package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingController(context *gin.Context) {
	context.Header("Content-Type", "text/html")
	context.String(http.StatusOK, "<h1>Hello World! from Builder Service</h1>")
}
