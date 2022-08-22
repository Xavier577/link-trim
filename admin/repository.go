package admin

import "gorm.io/gorm"

type AdminRepository struct {
	DbClient *gorm.DB
}

type IAdminRepository interface {
	FindAdminById(id uint) (Admin, error)
	FindAllAdmins() ([]Admin, error)
}

func (adminRepository *AdminRepository) FindAdminById(id uint) (Admin, error) {
	return Admin{}, nil
}

func (adminRepository *AdminRepository) FindAllAdmins() ([]Admin, error) {
	var admins []Admin

	err := adminRepository.DbClient.Find(&admins).Error

	return admins, err
}
