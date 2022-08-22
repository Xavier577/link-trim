package auth

import (
	"example/trim-server/models"
)

type LoginDto struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginSuccessResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type AuthFailureResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type SignUpSuccess struct {
	Message string `json:"message"`
	User    struct {
		UserID       string               `json:"user_id"`
		Username     string               `json:"username"`
		Email        string               `json:"email"`
		TrimmedLinks []models.TrimmedLink `json:"trimmed_links"`
	} `json:"user"`
}
