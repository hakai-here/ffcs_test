package migrations

import (
	"ffcs/api/db"
	"ffcs/pkg/models"
)

func Migrate() {
	database := db.Connect()
	database.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	database.AutoMigrate(&models.Users{}, &models.SelectedCourses{}, &models.Slots{})
}
