package session_repository

import "gorm.io/gorm"

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}
