package trimmedlinks

type CreateTrimmedLinkDto struct {
	LinkUrl    string `json:"link_url" binding:"required"`
	TrimmedUrl string
	UserId     uint
}
