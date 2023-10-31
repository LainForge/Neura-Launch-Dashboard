package initializers

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	DB, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to DB")
	} else {
		log.Info("Successfully connected to Database!")
	}
}
