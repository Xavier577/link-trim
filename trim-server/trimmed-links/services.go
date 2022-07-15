package trimmedlinks

import (
	"example/trim-server/database"
	"example/trim-server/global"
)

type TrimmedLinkService struct{}

var trimmedLinkRepository = new(TrimmedLinkRepository)

func (trimmedLinkService *TrimmedLinkService) FetchTrimmedLink() string {
	return "trimmed link"
}

func (trimmedLinkService *TrimmedLinkService) CreateTrimmedLink(createTrimmedLinkDto *CreateTrimmedLinkDto) (database.TrimmedLink, error) {
	trimmedLink, err := global.UUID()

	if err != nil {
		panic(err)
	}

	createTrimmedLinkDto.TrimmedUrl = trimmedLink
	return trimmedLinkRepository.CreatedTrimmedLink(createTrimmedLinkDto)
}
