package models

type TrimmedLink struct {
	ID      uint `gorm:"primaryKey"`
	Link    string
	Trimmed string
	UserId  uint
}

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	TrimmedLinks []TrimmedLink
}
