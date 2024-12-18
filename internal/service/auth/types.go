package auth_service

import (
	"github.com/kageyama0/chotto-rental/config"
	"github.com/kageyama0/chotto-rental/internal/model"
	session_repository "github.com/kageyama0/chotto-rental/internal/repository/session"
	user_repository "github.com/kageyama0/chotto-rental/internal/repository/user"
	"gorm.io/gorm"
)

type UserRepository = user_repository.UserRepository
type SessionRepository = session_repository.SessionRepository
type AuthConfig = config.AuthConfig

type Session = model.Session
type User = model.User
type DeviceInfo = model.DeviceInfo

type AuthService struct {
	db                *gorm.DB
	sessionRepository *SessionRepository
	userRepository    *UserRepository
	config            *AuthConfig
}
