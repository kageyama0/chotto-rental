package session_repository

import (
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"gorm.io/gorm"
)

func (r *SessionRepository) FindByID(id uuid.UUID) (*Session, int) {
	var session Session
	err := r.db.Where("id = ? AND is_valid = ?", id, true).First(&session).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, e.NOT_FOUND_SESSION
		}
		return nil, e.SERVER_ERROR
	}

	return &session, e.OK
}
