package user_repository

import "gorm.io/gorm"

// Craate: ユーザーを作成
func (r *UserRepository) Create(user *User) error {
	err := r.db.Create(user).Error
	return err
}

// CreateWithTransaction: トランザクション内でユーザーを作成
func (r *UserRepository) CreateWithTransaction(tx *gorm.DB, user *User) error {
	err := tx.Create(user).Error
	return err
}
