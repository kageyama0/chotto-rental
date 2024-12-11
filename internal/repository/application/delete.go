package application_repository

import (
	"github.com/google/uuid"
)

// -- DeleteByUserID: ユーザーIDを使用してアプリケーションを削除する
func (r *ApplicationRepository) DeleteByUserID(userID *uuid.UUID) error {
	err := r.db.Where("applicant_id = ?", userID).Delete(&Application{}).Error
	if err != nil {
		return err
	}

	return nil
}
