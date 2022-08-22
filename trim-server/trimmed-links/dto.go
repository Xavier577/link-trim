package trimmedlinks

type CreateTrimmedLinkDto struct {
	LinkUrl    string `json:"link_url" binding:"required"`
	TrimmedUrl string `swaggerignore:"true"`
	UserId     string `swaggerignore:"true"`
}
