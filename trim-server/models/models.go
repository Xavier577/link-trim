package models

import (
	"time"
)

type TrimmedLink struct {
	PK        uint `gorm:"primaryKey"`
	Link      string
	Trimmed   string
	UserId    string
	CreatedAt time.Time
	DeletedAt time.Time
}

type User struct {
	PK           uint   `gorm:"primaryKey"`
	UserId       string `gorm:"type:uuid;unique"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	Role         Role
	Admin        Admin
	TrimmedLinks []TrimmedLink
	CreatedAt    time.Time
	DeletedAt    time.Time
}

type Admin struct {
	PK     uint `gorm:"primaryKey"`
	UserId string
}

type Role struct {
	PK     uint `gorm:"primaryKey"`
	Type   string
	UserId string
}
