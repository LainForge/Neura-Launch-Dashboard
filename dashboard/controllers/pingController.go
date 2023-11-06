package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
    ctx.JSON(
        http.StatusOK,
        gin.H{
            "message": "Hello World!",
        })
}

func Hello(ctx *gin.Context) {
    ctx.Header("Content-Type", "text/html; charset=utf-8")
    ctx.String(http.StatusOK, "<h1>Hello World! Kunal is a fucking genius</h1>")
}
