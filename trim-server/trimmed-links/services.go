package trimmedlinks

func getLinkById(id uint, linkRepo ILinkRepository) (int64, error) {

	link := linkRepo.FindLink(id)

	return link.RowsAffected, link.Error
}
