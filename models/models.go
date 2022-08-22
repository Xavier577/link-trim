package models

import (
	"time"
)

type User struct {
	UserID       string `gorm:"primaryKey;type:uuid;"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	Role         Role
	Admin        Admin
	TrimmedLinks []TrimmedLink `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time
	DeletedAt    time.Time
}

type TrimmedLink struct {
	PK        uint `gorm:"primaryKey"`
	Link      string
	Trimmed   string
	UserID    string
	CreatedAt time.Time
	DeletedAt time.Time
}
type Admin struct {
	PK     uint `gorm:"primaryKey"`
	UserID string
}

type Role struct {
	PK     uint `gorm:"primaryKey"`
	Type   string
	UserID string
}
