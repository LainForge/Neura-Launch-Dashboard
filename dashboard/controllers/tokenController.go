package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/LainForge/Neura-Launch-Dashboard/dashboard/helpers"
    "golang.org/x/crypto/bcrypt"
)

func TokenController(ctx *gin.Context) {

    var body struct {
        Email string
    }

    err := ctx.Bind(&body)
    helpers.CheckErrorWithMessage(err, "Failed to read body")

    hash, err := bcrypt.GenerateFromPassword([]byte(body.Email), 10)
    helpers.CheckError(err)

    ctx.JSON(
        200,
        gin.H{"token": string(hash)},
    )

}
