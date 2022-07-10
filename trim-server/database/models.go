package database

type TrimmedLink struct {
	ID      uint `gorm:"primaryKey"`
	User    User
	UserId  string
	Link    string
	Trimmed string
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}
