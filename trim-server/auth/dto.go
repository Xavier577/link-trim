package auth

type LoginDto struct {
	UsernameOrEmail string `json:"username_or_email"  binding:"required"`
	Password        string `json:"password" binding:"required"`
}
