package controllers

import (
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func VerifyCheckSumControllers(context *gin.Context) {
	var body struct {
		checkSum string
		filePath string
	}

	err := context.Bind(&body)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(body.filePath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	hasher := sha256.New()
	evaluatedCheckSum, err := io.Copy(hasher, f)
	if err != nil {
		panic(err)
	}

	var result bool = string(evaluatedCheckSum) == body.checkSum

	context.JSON(
		http.StatusOK,
		gin.H{
			"result": result,
		},
	)
}
