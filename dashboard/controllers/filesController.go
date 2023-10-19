package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/LainForge/Neura-Launch-Dashboard/initializers"
	"github.com/LainForge/Neura-Launch-Dashboard/models"
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

	sess, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION), Credentials: credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"")})
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

	// Get a file from the form input name "file" and "checksum"
	file, header, err := c.Request.FormFile("file")
	checksum := c.Request.FormValue("checksum")

	// extract the project token from the file name which is token.zip
	token := strings.Split(header.Filename, ".")[0]

	// Check if the project exists
	var project models.Project
	result := initializers.DB.Where("token = ?", token).First(&project)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project does not exist"})
		return
	}

	// if the project exists then save the checksum
	project.Checksum = checksum
	initializers.DB.Save(&project)

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

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "bucket": AWS_S3_BUCKET})
}
