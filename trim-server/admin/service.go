package admin

type AdminService struct {
	adminRepository IAdminRepository
}

type IAdminService interface {
	GetAdmins() ([]Admin, error)
}

func (adminService *AdminService) GetAdmins() ([]Admin, error) {
	return adminService.adminRepository.FindAllAdmins()
}
