package orm

import (
	"fmt"

	"gorm.io/gorm"
)

var Driver *gorm.DB

func LoadDriver(db *gorm.DB) {
	Driver = db

	fmt.Println(Driver != nil)
}
