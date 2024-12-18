package auth_handler

import (
	auth_service "github.com/kageyama0/chotto-rental/internal/service/auth"
	"gorm.io/gorm"
)

func NewAuthHandler(db *gorm.DB, authService *auth_service.AuthService) *AuthHandler {
	return &AuthHandler{
		db:          db,
		authService: authService,
	}
}
