package users

import (
	"example/trim-server/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User models.User

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

	id, err := uuid.NewV1()

	if err != nil {
		return err
	}
	user.UserID = id.String()

	return
}
