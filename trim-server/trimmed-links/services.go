package trimmedlinks

import (
	"example/trim-server/shared/uuid"
	"log"
)

type TrimmedLinkService struct {
	trimmedLinkRepository ITrimmedLinkRepository
	uuidService           uuid.IUUIDService
}

type ITrimmedLinkService interface {
	FetchOriginalLink(trimmedLinkUUID string) (bool, TrimmedLink, error)
	CreateTrimmedLink(createTrimmedLinkDto *CreateTrimmedLinkDto) (TrimmedLink, error)
}

func (TLS *TrimmedLinkService) FetchOriginalLink(trimmedLinkUUID string) (bool, TrimmedLink, error) {
	return TLS.trimmedLinkRepository.FindOriginalLink(trimmedLinkUUID)
}

func (TLS *TrimmedLinkService) CreateTrimmedLink(createTrimmedLinkDto *CreateTrimmedLinkDto) (TrimmedLink, error) {
	trimmedLink, err := TLS.uuidService.Generate()

	if err != nil {
		log.Fatal(err.Error())

		return TrimmedLink{}, err
	}

	createTrimmedLinkDto.TrimmedUrl = trimmedLink

	return TLS.trimmedLinkRepository.CreateTrimmedLink(createTrimmedLinkDto)
}
