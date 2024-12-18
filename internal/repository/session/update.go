package session_repository

import "github.com/google/uuid"

// -- Invalidate : セッションを無効化する
func (r *SessionRepository) Invalidate(id uuid.UUID) error {
	return r.db.Model(&Session{}).Where("id = ?", id).Update("is_valid", false).Error
}
