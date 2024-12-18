package auth_service

import (
	"github.com/kageyama0/chotto-rental/config"
	session_repository "github.com/kageyama0/chotto-rental/internal/repository/session"
	user_repository "github.com/kageyama0/chotto-rental/internal/repository/user"
	"gorm.io/gorm"
)

func NewAuthService(db *gorm.DB, config *config.AuthConfig) *AuthService {
	userRepository := user_repository.NewUserRepository(db)
	sessionRepository := session_repository.NewSessionRepository(db)
	return &AuthService{
		db:                db,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		config:            config,
	}
}
