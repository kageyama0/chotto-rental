package session_repository

import (
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type Session = model.Session
type SessionRepository struct {
	db *gorm.DB
}
