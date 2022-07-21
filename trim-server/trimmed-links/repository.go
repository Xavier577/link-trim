package trimmedlinks

import (
	"errors"

	"gorm.io/gorm"
)

type TrimmedLinkRepository struct {
	dbClient *gorm.DB
}

type ITrimmedLinkRepository interface {
	CreateTrimmedLink(createLinkDto *CreateTrimmedLinkDto) (TrimmedLink, error)
	FindOriginalLink(trimmedLinkUUID string) (bool, TrimmedLink, error)
}

func (linkRepo *TrimmedLinkRepository) CreateTrimmedLink(createLinkDto *CreateTrimmedLinkDto) (TrimmedLink, error) {
	trimmedLink := TrimmedLink{UserId: createLinkDto.UserId,
		Link:    createLinkDto.LinkUrl,
		Trimmed: createLinkDto.TrimmedUrl}

	result := linkRepo.dbClient.Create(&trimmedLink)

	return trimmedLink, result.Error

}

func (linkRepo *TrimmedLinkRepository) FindOriginalLink(trimmedLinkUUID string) (bool, TrimmedLink, error) {
	var trimmedLink TrimmedLink
	err := linkRepo.dbClient.Select("link").First(&trimmedLink, &TrimmedLink{Trimmed: trimmedLinkUUID}).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, trimmedLink, err
}

// func (linkRepo *TrimmedLinkRepository) FindAllTrimmedLinks() {

// }

// func (linkRepo *TrimmedLinkRepository) FindUsersTrimmedLink(userId uint) {

// }
