package migrations

import (
	"hoangvuphone/internal/models"
	"gorm.io/gorm"
)


func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
	
}