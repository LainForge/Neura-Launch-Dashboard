package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	DB, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	} else {
		fmt.Print("Successfully connected to Database!")
	}
}
