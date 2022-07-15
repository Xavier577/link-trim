package database

type TrimmedLink struct {
	ID      uint `gorm:"primaryKey"`
	Link    string
	Trimmed string
	UserId  uint
}

type User struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Email        string
	Password     string
	TrimmedLinks []TrimmedLink
}
