package initializers

import "github.com/rohansingh9001/Neura-Launch-Dashboard/dashboard/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
