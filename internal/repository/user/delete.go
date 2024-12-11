package user_repository

import (
	"github.com/google/uuid"
)

// DeleteByID: IDでユーザーを削除
func (r *UserRepository) DeleteByID(id *uuid.UUID) error {
	err := r.db.Delete(&User{}, "id = ?", id).Error
	return err
}
