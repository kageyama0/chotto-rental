package auth_handler

import (
	"github.com/kageyama0/chotto-rental/internal/model"
	auth_service "github.com/kageyama0/chotto-rental/internal/service/auth"
	"gorm.io/gorm"
)


type User = model.User
type DeviceInfo = model.DeviceInfo
type AuthHandler struct {
	db          *gorm.DB
	authService *auth_service.AuthService
}

// @Description ユーザー登録リクエスト
type SignupRequest struct {
	DisplayName string `json:"displayName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
}

// @Description ログインリクエスト
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
