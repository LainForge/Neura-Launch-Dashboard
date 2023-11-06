package controllers

import (
    "crypto/sha256"
    "io"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
)

func VerifyCheckSumController(context *gin.Context) {
    var body struct {
        checkSum string
	filePath string
    }

    err := context.Bind(&body)
    if err != nil {
        log.Panic(err)
    }

    f, err := os.Open(body.filePath)
    if err != nil {
        log.Panic(err)
    }

    defer f.Close()

    hasher := sha256.New()
    evaluatedCheckSum, err := io.Copy(hasher, f)
    if err != nil {
        log.Panic(err)
    }

    var result bool = string(evaluatedCheckSum) == body.checkSum

    context.JSON(
        http.StatusOK,
        gin.H{
            "result": result,
	},
    )
}
