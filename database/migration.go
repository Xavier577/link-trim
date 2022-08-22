package database

import (
	"example/trim-server/models"

	"gorm.io/gorm"
)

func LoadMigrations(dbDriver *gorm.DB) {
	dbDriver.AutoMigrate(&models.User{})
	dbDriver.AutoMigrate(&models.Role{})
	dbDriver.AutoMigrate(&models.TrimmedLink{})
}
