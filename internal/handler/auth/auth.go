package auth_handler

import (
	"github.com/kageyama0/chotto-rental/pkg/auth"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db          *gorm.DB
	authService *auth.AuthService
}

func NewAuthHandler(db *gorm.DB, authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		db:          db,
		authService: authService,
	}
}
