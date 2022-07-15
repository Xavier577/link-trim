package trimmedlinks

import (
	"errors"
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

type TrimmedLinkRepository struct{}

func (linkRepo *TrimmedLinkRepository) CreatedTrimmedLink(createLinkDto *CreateTrimmedLinkDto) (database.TrimmedLink, error) {
	trimmedLink := database.TrimmedLink{UserId: createLinkDto.UserId,
		Link:    createLinkDto.LinkUrl,
		Trimmed: createLinkDto.TrimmedUrl}
	result := db.Create(&trimmedLink)

	return trimmedLink, result.Error

}

func (linkRepo *TrimmedLinkRepository) FindTrimmedLink(id uint) (bool, database.TrimmedLink, error) {
	var trimmedLink database.TrimmedLink
	result := db.First(trimmedLink, &database.TrimmedLink{ID: id})
	isNotFoundErr := errors.Is(result.Error, gorm.ErrRecordNotFound)
	err := result.Error
	return isNotFoundErr, trimmedLink, err
}

// func (linkRepo *TrimmedLinkRepository) FindAllTrimmedLinks() {

// }

// func (linkRepo *TrimmedLinkRepository) FindUsersTrimmedLink(userId uint) {

// }
