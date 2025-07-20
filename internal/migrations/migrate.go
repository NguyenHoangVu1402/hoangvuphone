package migrations

import (
	"gorm.io/gorm"
	"hoangvuphone/internal/models"
)

func MigrateDatabase(db *gorm.DB) error {
	tables := []interface{}{
		&models.Permission{},
		&models.Role{},
		&models.Account{},
		&models.RolePermissions{},
		&models.AccountRoles{},
	}

	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			return err
		}
	}
	return nil
}