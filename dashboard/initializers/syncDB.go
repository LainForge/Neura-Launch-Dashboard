package initializers

import "github.com/LainForge/Neura-Launch-Dashboard/dashboard/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
