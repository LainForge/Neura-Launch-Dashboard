package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

const (
	AWS_S3_REGION = "ap-southeast-1"
	AWS_S3_BUCKET = "neura-launch-zips"
)

func connectAWS() *session.Session {

	fmt.Println("Connecting to AWS", os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"))

	sess, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION),Credentials: credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		""),})
	if err != nil {
		panic(err)
	}
	return sess
}

func UploadFile(c *gin.Context) {

	var sess = connectAWS()

	// Parse the multipart form, with max file size 100MB
	err := c.Request.ParseMultipartForm(100 << 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Get a file from the form input name "file"
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve the file from the form"})
		return
	}
	defer file.Close()

	filename := header.Filename

	// Upload the file to S3.
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET), // Bucket
		Key:    aws.String(filename),      // Name of the file to be saved
		Body:   file,                      // File
	})

	fmt.Println("Error........: ", err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "bucket": AWS_S3_BUCKET})
}
