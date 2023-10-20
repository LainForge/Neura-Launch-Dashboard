package controllers

import (
	"context"
	"fmt"
	"os"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func MakeDockerImage(ctx *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		 ctx.String(500, "Error creating Docker Client")
		 return
	}

	// Build an image from a Dockerfile in the current directory
	buildcontext, err := os.Open("./model.tar")
	if err != nil {
		ctx.String(500, "Error in reading file ", err)
		return 
	}

	// Building an image from the Dockerfile
	buildResponse, err :=  cli.ImageBuild( 
		context.Background(),
		buildcontext,
		types.ImageBuildOptions{
            Dockerfile: "model/Dockerfile",
            Tags:       []string{"my-image:latest"},
        },
	)
	if err != nil {
		ctx.String(500, "Error in Making Image %v", err)
		return
	}

	// Reading the output of the build 

	// Making a buffer slice of size 1KB
	buffer := make([]byte, 1024)

	// Iterating through the build Response body
	for {
		n, err := buildResponse.Body.Read(buffer)
		if n > 0 {
            fmt.Print(string(buffer[:n]))
        }
        if err != nil {
			ctx.String(500, "There was an error in building Image")
            break
        }
	}
}