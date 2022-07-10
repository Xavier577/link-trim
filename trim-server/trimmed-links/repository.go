package trimmedlinks

import (
	"example/trim-server/database"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeLinkRepository(db_driver *gorm.DB) {
	db = db_driver
	error := db_driver.AutoMigrate(&database.TrimmedLink{})

	if error != nil {
		panic(error)
	}
}

type LinkRepositoryStruct struct{}

type ILinkRepository interface {
	FindLink(id uint) *gorm.DB
}

func (linkRepo *LinkRepositoryStruct) FindLink(id uint) *gorm.DB {
	result := db.First(&database.TrimmedLink{ID: id})

	return result
}
