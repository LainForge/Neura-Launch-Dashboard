package controllers

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type AWSConfig struct {
	AWSAccessKeyID string
	AWSSecretAccessKey string 
	AWSS3Region string
	AWSS3Bucket string
	AWSS3Client *s3.S3
}

func (config AWSConfig) ConnectToAWS() {
	// Creating a new session
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(config.AWSS3Region),
			Credentials: credentials.NewStaticCredentials(config.AWSAccessKeyID, config.AWSSecretAccessKey, ""),
		},
	)

	if err != nil {
		return
	}

	config.AWSS3Client = s3.New(sess)
}

func (config AWSConfig) DownloadFilesFromS3() error {
	// Defining Global S3 Client 
	s3Client := config.AWSS3Client

	// Getting the object to download 
	downloader := s3.GetObjectInput{
		Bucket: aws.String(config.AWSS3Bucket),
		Key: aws.String("test.txt"),
	}

	// Downloading the object
	result, err := s3Client.GetObject(&downloader)

	if err != nil {
		return err
	}

	// Creating a Zip file 
	zipFile, err := os.Create("test.zip")
	if err != nil {
		return err
	}

	// Copying the object to the zip file
	_, err = io.Copy(zipFile, result.Body)
	if err != nil {
		return err
	}

	return nil
	
}


func downloadFilesController(context *gin.Context) {

}
