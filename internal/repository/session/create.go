package session_repository

import "gorm.io/gorm"

// Create : セッション情報をDBに保存する
func (r *SessionRepository) Create(session *Session) error {
	err := r.db.Create(session).Error
	return err
}

// CreateWithTransaction : トランザクション内でセッション情報をDBに保存する
func (r *SessionRepository) CreateWithTransaction(tx *gorm.DB, session *Session) error {
	err := tx.Create(session).Error
	return err
}
