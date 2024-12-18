package user_repository

import (
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"gorm.io/gorm"
)

// FindByID: IDでユーザーを取得
func (r *UserRepository) FindByID(id uuid.UUID) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}


// FindByEmail: メールアドレスでユーザーを取得
func (r *UserRepository) FindByEmail(email string) (*User, int) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
					return nil, e.NOT_FOUND_USER
			}
			return nil, e.SERVER_ERROR
	}
	return &user, e.OK
}
