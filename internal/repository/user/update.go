package user_repository

import (
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"gorm.io/gorm"
)

// Update: ユーザー情報を更新
func (r *UserRepository) Update(id uuid.UUID, user *User) int {
	err := r.db.Model(&User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return e.NOT_FOUND_USER
		}
		return e.SERVER_ERROR
	}
	return e.OK
}
