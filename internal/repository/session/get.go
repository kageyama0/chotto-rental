package session_repository

import "github.com/google/uuid"

func (r *SessionRepository) FindByID(id uuid.UUID) (*Session, error) {
	var session Session
	err := r.db.Where("id = ? AND is_valid = ?", id, true).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}
