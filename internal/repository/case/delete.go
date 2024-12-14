package case_repository

import (
	"github.com/google/uuid"
)

// - DeleteByID: idで案件を削除
func (r *CaseRepository) DeleteByID(id uuid.UUID) error {
	err := r.db.Delete(&Case{}, "id = ?", id).Error
	return err
}

// - DeleteByUserID: ユーザーIDを使用して案件を削除する
func (r *CaseRepository) DeleteByUserID(userID *uuid.UUID) error {
	err := r.db.Where("user_id = ?", userID).Delete(&Case{}).Error
	return err
}
